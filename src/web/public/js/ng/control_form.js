angular.module('app').controller('ControlForm', ['$scope', '$tank', function ( $scope, $tank ) {

  $scope.$watch( function(){return $tank.LeftMotor}, function(val){ $scope.LeftMotor = $tank.LeftMotor } )
  $scope.$watch( function(){return $tank.RightMotor}, function(val){ $scope.RightMotor = $tank.RightMotor } )

  $scope.apply = function () {

    $tank.LeftMotor = $scope.LeftMotor
    $tank.RightMotor = $scope.RightMotor

    $scope.sendTankCommand({
      LeftMotor: parseFloat($scope.LeftMotor),
      RightMotor: parseFloat($scope.RightMotor)
    })
  }

  $scope.stopMotors = function () {
    $scope.LeftMotor = 0
    $scope.RightMotor = 0

    $scope.apply()
  }

  $scope.fire = function () {
    $scope.sendTankCommand({ Fire: true })
  }
}]);
