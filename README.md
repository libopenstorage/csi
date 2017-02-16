# Specifications for Container Storage Interfaces

The purpose of this project is to define the various (vendor agnostic) interfaces between cloud native schedulers and persistent data services.

### Issues addressed by this spec

1. Deployment of the data service provider by the scheduler: The data service provider, perhaps packaged as a container should be deployed on the servers being managed by the scheduler.

2. Inline storage service provisioning: Users should be able to allocate the data service resources programatically via the scheduler interface. This obsoletes the need to do static out-of-band volume provisioning.

3. Data locality aware scheduling: The scheduler should be able to take into account the locality of a container's data, before it schedules it on a host.

4. Scheduler driven data life cycle management: The life cycle of a container and it's storage are different. The scheduler should be able to manage both, separately as independent entities.

5. Propagation of the data service properties via the scheduler: When a storage resource is created, the properties of the resource should be transparently passed through by the scheduler. This obsoletes the need for such information being provided out-of-band via other methods.

6. Common protocol of communication via the data service, scheduler and container runtime engine: The data service provider should be able to allocate resources and manage them using the same (or close to similar) protocol regardless of it being used by the scheduler agent or the container runtime engine, like Docker.

7. Application awareness facilitated to the data service layer: The data service layer should have broader context of the application that is being deployed. An application will comprise of many containers and having a broader context enables the data service layer to optimize and implement the correct HA features. As an example, consider the deployment of a Cassandra ring. Knowing the various containers that are part of the ring will help the data service provider to appropriately place the data in the correct failure domains.

8. Authentication of access to a data service facilitated via the scheduler: Prohibit a data service provider from allowing a container to use a service (such as a volume) that it is not authorized to use.

9. Monitoring - Alerts and Stats propagated via the scheduler: A common framework to get alerts and stats via the scheduler is desired. This prevents the need for external event correlation.

### Organization of the spec
This spec covers two aspects of orchestrating the deployment of data services via a scheduler:

1. The bootstrap deployment of the data service container itself.
2. The runtime communication between a scheduler agent and the data service container.

## Bootstrap Deployment of Data Service Resources
This section of the spec describes how data service providers are deployed by orchestration software.  For example, these providers can be packaged as Linux Containers and they would need to be depoyed on the physical infrastructure by the orchestration software.  This is specified in [api/bootstrap.go](api/bootstrap.go).

## Runtime communication between the scheduler and the data service
Once the data service has been deployed, there are 4 specific interfaces that schedulers and data service providers need to implement.  This is specified in [api/provider.go](api/provider.go).  The scheduler and the provider could communicate via a runtime `UNIX sock` file on the agent host machine (TBD).

### 1. Discovery of Data Services
Applications that rely on data services should be able to dynamically discover where the provisioned resources are available.  The data service API should also be able to influence where and when these services should be scheduled based on the underlying constraints.

### 2. Provisioning and Instantiation of Data Services
The allocation, use (read and write) and destruction (what used to be known as CRUD) needs to be orchestrated through this interface.

### 3. Lifecycle Operations on Data Services (TBD)
Data state and its lifecycle, such as retention levels, version levels, access controls should be separated from the actual application that uses them.  It should also be controlled by the scheduling software and it is the goal of this API to define how that is goverened.

### 4. Security (TBD)
This defines a set of constraints around how a container can authenticate itself in order to operate on a storage service.  This would prevent a container launched by a user from accessing a volume they do not have access to.  

## Licensing
`CNCF-CSI` is licensed under the Apache License, Version 2.0. See LICENSE for the full license text.

## Contributing
Want to collaborate and add? Here are instructions to [get started contributing code](contributing.md)
