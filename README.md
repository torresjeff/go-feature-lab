# Feature Lab (Go feature flags)

Feature Lab is a feature flag solution developed in Go (Golang). 
Feature flags are a powerful technique that allows you to turn features on and off, or to serve different versions of features to different users.
This allows you to test and experiment new features in a controlled and safe manner.

## Treatment allocation
When using Feature Lab, it is important to ensure that you're allocating users to different treatments in a fair and consistent manner.
Treatment allocation is the process of assigning users to different treatments based on a set of criteria.
For example, you may want to assign different treatments to users based on their geographic location, device type, user ID, session ID, etc.

You should strive for your allocation criteria to be deterministic: instead of using the session ID as your allocation criteria
(which changes everytime the user logs out), you should use the user ID instead (doesn't change between sessions).
Using a value that changes over time as your allocation criteria could have the effect of a user being assigned to different treatments between sessions (or in the same session if the value can change multiple times during the session),
resulting in an inconsistent experience.

## How Feature Lab works
The basic building block of Feature Lab is a `Feature`. A `Feature` consists of a name and 0 or more `FeatureAllocations`.
Suppose we have a live-streaming website, and we're working on a new feature to show recommended channels to a user based on the content that they watch.
Our feature name would be "Show recommendations" and we're going to define different treatments to decide the best placement of the new recommendations panel inside the webpage.

A `FeatureAllocation` represents the weight that is given to a specific treatment (eg. "Control", "Treatment 1", etc.).
In the context of experiments, weights are used to assign a proportion of users to each treatment.
Weights are typically expressed as a percentage or a fraction, and they represent the relative size of each treatment group.

For example, here are the weights for our new recommendation feature:

| **Name** | **Weight** |
|----------|------------|
| C        | 30         |
| T1       | 50         |
| T2       | 20         |
| **Total**| 100        |
Note: in this table C = Control Treatment, T1 = Treatment 1, T2 = Treatment 2.

The actual meaning of each treatment is given entirely by you. Here's the meaning I've assigned to each treatment:
* Control treatment (`C`): recommendations are not shown to the user. The control treatment is what we're benchmarking against.
* Treatment 1 (`T1`): recommendations are shown at the top of the page.
* Treatment 2 (`T2`): recommendations are shown at the right-side of the page.

According to the weights shown in the table, the probability that a user is assigned treatment `C` is 30%; `T1` is 50%; and `T2` is 20%.

Note that weights need not add up to 100. The relative weight of each treatment is determined by dividing the weight of that treatment by the total weight of all treatments.

**Note:** if you want to gather data to benchmark against your control treatment fairly, then the probability of a user falling in each treatment should be equal (eg. weights could be: C - 10; T1 - 10; T2 - 10).
This way you'll have a uniform distribution of users for each treatment.


## Examples
See `cmd/gettreatment/main.go`

## TODO list
- [ ] Feature Lab daemon that syncs allocation data every 10 minutes
  - [ ] Cache allocation data for faster queries
  - [ ] gRPC support for querying treatments
  - [ ] Can be deployed as a sidecar with your main app container
- [ ] Report metrics about feature triggers