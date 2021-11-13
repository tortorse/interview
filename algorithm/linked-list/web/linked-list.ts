class List {
  start: List | null;
  end: List | null;
  data: string;

  public add(data: string) {
    const list = new List()
    list.data = data
    if (!this.start) {
      this.start = list
    }
    list.start = this

    if (!this.end) {
      this.end = list
    } else {
      this.end.start = list
      list.end = this.end
      this.end = list
    }
  }
}




