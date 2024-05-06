const process = require("process");
const pre = process.memoryUsage();
const re = /abc/g;
const post = process.memoryUsage();
console.log("used", post.heapUsed, post.heapUsed - pre.heapUsed);
console.log("total", post.heapTotal, post.heapTotal - pre.heapTotal);
console.log("external", post.external, post.external - pre.external);
