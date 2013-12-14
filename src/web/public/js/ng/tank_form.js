angular.module('app').controller('TankForm', ['$scope', '$tank', function ( $scope, $tank ) {
  $scope.$watch( function(){return $tank.Id}, function(val){ $scope.Id = $tank.Id } )
  $scope.$watch( function(){return $tank.Health}, function(val){ $scope.Health = $tank.Health } )

  $scope.$watch( function(){return $tank.Coords.X}, function(val){ $scope.X = $tank.Coords.X } )
  $scope.$watch( function(){return $tank.Coords.Y}, function(val){ $scope.Y = $tank.Coords.Y } )

  $scope.$watch( function(){return $tank.LeftMotor}, function(val){ $scope.LeftMotor = $tank.LeftMotor } )
  $scope.$watch( function(){return $tank.RightMotor}, function(val){ $scope.RightMotor = $tank.RightMotor } )

  $scope.$watch( function(){return $tank.Direction}, function(val){ $scope.Direction = $tank.Direction } )

  $scope.$watch( function(){return $tank.Gun.Direction}, function(val){ $scope.GunDirection = $tank.Gun.Direction } )
  $scope.$watch( function(){return $tank.Gun.Temperature}, function(val){ $scope.GunTemperature = $tank.Gun.Temperature } )
  $scope.$watch( function(){return $tank.Gun.ReloadProgress}, function(val){ $scope.GunReloadProgress = $tank.Gun.ReloadProgress } )


  $scope.$on("tank:message", function(event, message){
    $scope.setTank(message.Tank)
  })

  $scope.$on("world:message", function(event, message){
    tanks = message.Tanks

    i = _.findIndex(tanks, {Id: $tank.Id})
    if ( i >= 0 ) { $scope.setTank(tanks[i]) }

    $scope.$digest()
  })

  $scope.setTank = function ( t ) {
    $tank.Id = t.Id
    $tank.Health = t.Health
    $tank.RightMotor = t.RightMotor
    $tank.LeftMotor = t.LeftMotor
    $tank.Direction = t.Direction
    $tank.Coords.X = t.Coords.X
    $tank.Coords.Y = t.Coords.Y
    $tank.Gun.Direction = t.Gun.Direction
    $tank.Gun.ReloadProgress = t.Gun.ReloadProgress
    $tank.Gun.Temperature = t.Gun.Temperature
  }

  $scope.apply = function () {}
}])
