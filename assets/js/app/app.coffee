# Hack to get bootstrap working
window.jQuery = jQuery = require 'jquery'
Bootstrap  = require 'bootstrap'

React      = require "react"
dispatcher = require "./dispatcher/dispatcher.coffee"

Routes     = require './routes.cjsx'
