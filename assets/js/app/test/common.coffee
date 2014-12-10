chai = require("chai")

sinon = require("sinon")
sinonChai = require("sinon-chai")
sinonAsPromised = require("sinon-as-promised")
chaiAsPromised = require ("chai-as-promised")

chai.use sinonChai
chai.use chaiAsPromised
global.sinon = sinon
global.expect = chai.expect
