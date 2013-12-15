angular.module('app').controller('Main', function ( $scope ) {

  $scope.isPage = function (name) { return (window.location.hash || "#") == ("#" + name) }
  $scope.changePage = function () { setTimeout(function(){$scope.$digest()}, 0) }
})
