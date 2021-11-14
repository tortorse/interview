export class Node {
  next
  data

  constructor(data) {
    this.data = data
    this.next = null
  }
}

export class List {
  start = null; // 头结点
  end = null; // 尾节点

  add(data) {
    const node = new Node(data)
    if (!this.start) {
      this.start = node
    }

    if (this.end) {
      this.end.next = node
    }

    this.end = node
  }

  delete(val) {
    let cur = this.start
    let pre
    while(cur) {
      if (cur.data === val) {
        // 单节点
        if (this.start === this.end) {
          this.start = null
          this.end = null
          break
        }

        if (pre) {
          pre.next = cur.next
        } else {
          // 删第一个节点，处理start
          if (cur.next) {
            this.start = cur.next
          }
        }

        // 处理最后一个节点
        if (!cur.next && pre) {
          this.end = pre
        }
        break
      }
      pre = cur
      cur = cur.next
    }
  }

  insertAfter(anchor, val) {
    let cur = this.start
    while(cur) {
      if (cur.data === anchor) {
        const node = new Node(val)
        if (cur.next) {
          node.next = cur.next
        } else {
          this.end = node
        }
        cur.next = node
        break
      }
      cur = cur.next
    }
  }

  item(index) {
    let i = 0
    let cur = this.start
    while(cur) {
      i++
      if (i === index) {
        return cur
      }
      cur = cur.next
    }
    return null
  }

  each(cb) {
    let cur = this.start
    while(cur) {
      cb(cur)
      cur = cur.next
    }
  }

  reverse() {
    if (!this.start) {
      return
    }
    const self = this
    function travere(node, pre = null) {
      if (node.next) {
        travere(node.next, node)
      } else {
        self.start = node
      }
      node.next = pre
      if (!pre) {
        self.end = node
      }
    }
    travere(this.start)
  }

  circle() {
    let values = {}
    let hasCircle = false
    this.each((node) => {
      if (values[node.data]) {
        hasCircle = true
      }
      values[node.data] = true
    })
    return hasCircle
  }
}

window.List = List





