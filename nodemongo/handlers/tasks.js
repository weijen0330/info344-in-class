"use strict";

const express = require('express');
const Task = require('../models/tasks/task.js');

//export a function from this module 
//that accepts a tasks store implementation
module.exports = function(store) {
    // Creating a new Mux (like in Go)
    let router = express.Router();

    router.get('/v1/tasks', async (req, res, next) => {
        // try {
        //     let tasks = await store.getAll();
        //     res.json(tasks);
        // } catch(err) {
        //     // jump to the next error handler in the chain
        //     next(err);
        // }

        // equivalent
        store.getAll()
            .then(tasks => {
                res.json(tasks);
            })
            .catch(next);
    });

    router.post('/v1/tasks', async(req, res, next) => {
        try {
                                // this will be undefined if body-parser is not being used!
            let task = new Task(req.body);
            let err = task.validate();
            // We have two blocks here because this first one is for the validate
            // We can throw the error in the validate method, but that will throw a 500 instead of a 400, the same
            // as MongoDB connection error, which would be confusing
            if (err) {
                res.status(400).send(err.message);
            } else {
                let result = await store.insert(task);
                res.json(task);
            }
        } catch(err) {
            next(err);
        }
    });

    router.patch('/v1/tasks/:taskID', async (req, res, next) => {
       let taskID = req.params.taskID;
       try {
           let updatedTask = await store.setComplete(taskId, req.body.complete);
           res.json(updatedTask);
       } catch(err) {
           next(err);
       }

    });

    router.delete('/v1/tasks/:taskID', async (req, res, next) => {
       let taskID = req.params.taskID;
       try {
           await store.delete(taskID);
           res.send(`deleted task ${taskID}`);
       } catch (err) {
           next(err)
       }
    });

    return router;
};
