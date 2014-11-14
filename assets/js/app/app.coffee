# Hack to get bootstrap working
window.jQuery = jQuery = require 'jquery'

React      = require "react"
kootstrap  = require "bootstrap"
dispatcher = require "./dispatcher/dispatcher.coffee"

Nav        = require './cjsx/nav.cjsx'
Hud        = require './cjsx/hud.cjsx'
