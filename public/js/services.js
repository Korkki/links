(function(){
    "use strict";
    angular.module("links")
        .service('alerts', Alerts)
        .service('links', Links);

    function Alerts() {
        var that = this;
        this.list = [];
        this.addAlert = addAlert;
        this.closeAlert = closeAlert;

        function addAlert(type, msg) {
            that.list.push({type: type, msg: msg});
        }

        function closeAlert(index) {
            that.list.splice(index, 1);
        }
    }

    Links.$inject = ['Restangular', 'alerts'];
    function Links(Restangular, alerts) {
        var links = Restangular.all("links");
        this.links = [];
        this.getLinks = getLinks;
        this.removeLink = removeLink;
        this.addLink = addLink;

        function getLinks (params) {
            var that = this, p = _.isUndefined(params) ? {} : params;
            links.getList(p).then(function(data){
                that.links = data;
            }, function(){
                alerts.addAlert({type: "danger", msg: "Error getting links!"});
            });
        };

        function removeLink(i, id){
            var that = this;
            links.one("", id).remove().then(function(){
                that.links.splice(i, 1);
                alerts.addAlert("success", "Link removed!");
            }, function(){
                alerts.addAlert("danger", "Error removing link!");
            });
        };

        function addLink(link){
            links.post(link).then(function(data){
                alerts.addAlert("success", "Link added!");
                return data;
            }, function(){
                alerts.addAlert("danger", "Error adding link!");
            });
        }
    }
})();
