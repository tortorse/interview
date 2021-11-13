var List = /** @class */ (function () {
    function List() {
    }
    List.prototype.add = function (data) {
        var list = new List();
        list.data = data;
        if (!this.start) {
            this.start = list;
        }
        list.start = this;
        if (!this.end) {
            this.end = list;
        }
        else {
            this.end.start = list;
            list.end = this.end;
            this.end = list;
        }
    };
    return List;
}());
