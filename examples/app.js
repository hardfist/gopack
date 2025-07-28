// Example JavaScript file for testing bundling
console.log('Hello from Gopack!');

// Import example (this would be parsed by a real bundler)
// import { helper } from './utils';

function greet(name) {
    return `Hello, ${name}!`;
}

function main() {
    const message = greet('World');
    console.log(message);
}

main();
