(function(){
    "use strict";
    angular.module("links")
        .controller("MainController", MainController)
        .controller("LinkModalFormController", LinkModalFormController);

    MainController.$inject = ['$modal', 'links'];
    function MainController($modal, links) {
        var vm = this;
        vm.links = [];
        vm.getLinks = links.getLinks;
        vm.removeLink = links.removeLink;
        vm.openLinkFormModal = openLinkFormModal;

        links.getLinks.apply(this);

        function openLinkFormModal(size){
            var modalInstance = $modal.open({
                templateUrl: "/partials/linkModalForm.html",
                controller: "LinkModalFormController",
                size: size
            });
        };
    }

    LinkModalFormController.$inject = ['$scope', 'Restangular', '$modalInstance', 'links'];
    function LinkModalFormController($scope, Restangular, $modalInstance, links) {
        $scope.link = {tags:[{}]};
        $scope.addLink = addLink;
        $scope.addTag = addTag;

        function addLink(valid){
            if (valid) {
                links.addLink($scope.link);
                $scope.link = {tags:[{}]};
            }
        }

        function addTag(){
            $scope.link.tags.push({});
        }
    }
})();
