function ControlForm( $scope ) {
  $scope.LeftMotor = 0
  $scope.RightMotor = 0

  $scope.apply = function () {
    message = {
      Type: "TankCommand",
      LeftMotor: parseFloat($scope.LeftMotor),
      RightMotor: parseFloat($scope.RightMotor)
    }

    $scope.sendMessage(message)
  }

  $scope.stopMotors = function () {
    $scope.LeftMotor = 0 
    $scope.RightMotor = 0

    $scope.apply()
  }
}
