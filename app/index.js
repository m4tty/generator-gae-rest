'use strict';
var util = require('util');
var path = require('path');
var yeoman = require('yeoman-generator');


var GaeRestGenerator = module.exports = function GaeRestGenerator(args, options, config) {
  yeoman.generators.Base.apply(this, arguments);

  this.on('end', function () {
    this.installDependencies({ skipInstall: options['skip-install'] });
  });

  this.pkg = JSON.parse(this.readFileAsString(path.join(__dirname, '../package.json')));
};

util.inherits(GaeRestGenerator, yeoman.generators.Base);

GaeRestGenerator.prototype.askFor = function askFor() {
  var cb = this.async();

  // have Yeoman greet the user.
  console.log(this.yeoman);

  var prompts = [{
    name: 'appName',
    message: 'What would you like to name your rest application?',
  }];

  this.prompt(prompts, function (props) {
    this.appName = props.appName;

    cb();
  }.bind(this));
};

GaeRestGenerator.prototype.app = function app() {
  this.mkdir('data');
  this.mkdir('domain');

  this.mkdir('web');
  this.mkdir('web/handlers');
  this.mkdir('web/resources');
  this.mkdir('web/static');

  
  this.copy('_package.json', 'package.json');
  this.copy('_bower.json', 'bower.json');
};

GaeRestGenerator.prototype.projectfiles = function projectfiles() {
  this.copy('editorconfig', '.editorconfig');
  this.copy('jshintrc', '.jshintrc');
};
