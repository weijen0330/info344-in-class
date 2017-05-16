"use strict";

// Loads modules from the node_module dir

// Gets back an Express function
const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const zips = require('./zips.json');
console.log('loaded %d zips', zips.length);

const zipCityIndex = zips.reduce((index, record) => {
    let cityLower = record.city.toLowerCase();
    let zipsForCity = index[cityLower];
    if (!zipsForCity) {
        index[cityLower] = zipsForCity = [];
    }
    zipsForCity.push(record);
    return index;
}, {});
console.log('there are %d zips in Seattle', zipCityIndex.seattle.length);
// Creates new Express application
const app = express();

const port = process.env.PORT || 80 ;
const host = process.env.HOST || 'localhost';

// Both Morgan and CORS returns middleware functions
// .use() --> every request needs to go through this middleware
app.use(morgan('dev')); // request logging
app.use(cors()); // adds access-control-allow-origin --> allows other people to call our API

app.get('/zips/city/:cityName', (req, res) => {
    let zipsForCity = zipCityIndex[req.params.cityName];
    if (!zipForCity) {
        res.status(404).send('invalid city name');
    } else {
        // .json will set content-type for you and parse it as JSON
        res.json(zipsForCity);
    }
});

// colon --> wildcard
// Whatever comes after the second slash, stores it in the variable name
// res object is equivalent to ResponseWriter in Go
app.get('/hello/:name', (req, res) => {
    res.send(`Hello ${req.params.name}!`);


});

// cool, template syntax
app.listen(port, host, () => {
    console.log(`server is listening at http://${host}:${port}...`);
});