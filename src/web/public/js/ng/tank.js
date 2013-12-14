angular.module('app').service('$tank', function () {

  this.Id = 0
  this.Health = 0
  this.Coords = { X: 0, Y: 0 }
  this.Direction = 0
  this.LeftMotor = 0
  this.RightMotor = 0
  this.Gun = {
    Direction: 0
  }

});
