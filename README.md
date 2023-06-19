
<p align="center">
    <img src="./assets/image/logo.svg" alt="Vince Logo" />
</p>

# vince

The Cloud Native Web Analytics Platform. Built on a persistent and fast key
value store that support ACID transactions with serializable snapshot isolation
(SSI) guarantees.



# Features

- Programmable alerting (using Typescript or modern Javascript ECMAScript 2022)
- Time on site tracking
- Conversion tracking 
- Multiple site management
- User Interaction Management 
- Campaign Management 
- Report Generation
- Goal Tracking 
- Event Tracking 
- Cloud Native (seamless k8s integration)
- Automatic TLS
- API for stats and sites management
- No runtime dependency (Static binary with everything you need)

## Usage

[Documentation](https://vinceanalytics.github.io/guide)

## Contributing

[Contributing](https://vinceanalytics.github.io/contibuting)

## Credits

All credits goes to [@plausible/analytics](https://github.com/plausible/analytics) . I am not affiliated with the project, if you love vince and want a managed version I recommend you try them out.


# Origins

This started as a go port of [Plausible](https://github.com/plausible/analytics), with 
the intention to remove clickhouse and postgresql dependency . **Vince** is
built on `sqlite` for operational data and [badger](https://github.com/dgraph-io/badger)
for timeseries.

The frontend is built with [@primer](https://github.com/primer) and web components.


# The UI

You might be wondering why the UI looks like github, we are using the same components
library as github [@primer](https://github.com/primer)

# The name vince 

Vince is named after [Vince Staples](https://en.wikipedia.org/wiki/Vince_Staples) , 
my favorite hip hop artist.

It was late night, 1 year after I quit my job and took a sabbatical, I was listening
to one of his album [Big Fish Theory](https://en.wikipedia.org/wiki/Big_Fish_Theory)
. The song Big Fish inspired me to program again.

The lyrics that got me going.
```
I was up late night ballin'
Countin' up hundreds by the thousand
```

So, enjoy the outcome of my late nights, hoping to be counting hundreds by the thousand
soon.
