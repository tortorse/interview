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

  deleteReverseN(n) {
    if (!this.start) {
      return
    }

    let i = 0
    function traverse(node, pre) {
      if (node.next) {
        traverse(node.next, node)
      }
      i++
      if (i === n) {
        if (node.next && pre) {
          pre.next = node.next
        }
      }
    }
    traverse(this.start)
  }

  middle() {
    if (!this.start) {
      return
    }

    let i = 1
    let j = 1
    let middle = []
    function traverse(node) {
      if (node.next) {
        i++
        traverse(node.next, node)
      }

      if (j * 2 === i) {
        middle.unshift(node.data)
      }

      if ((j - 1) * 2 === i) {
        middle.unshift(node.data)
      }

      if ((j - 1) * 2 + 1 === i) {
        middle.unshift(node.data)
      }
      j++
    }
    traverse(this.start)
    return middle
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

  toArr() {
    let arr = []
    this.each((node) => {
      arr.push(node.data)
    })
    return arr
  }

  concat(list) {
    const res = new List()
    let cur1 = this.start
    let cur2 = list.start

    while (cur1 || cur2) {
      if (!cur1) {
        res.add(cur2.data)
        cur2 = cur2.next
        continue
      }

      if (!cur2) {
        res.add(cur1.data)
        cur1 = cur1.next
        continue
      }

      if (cur1.data < cur2.data) {
        res.add(cur1.data)
        cur1 = cur1.next
      } else {
        res.add(cur2.data)
        cur2 = cur2.next
      }
    }

    console.log(res)
    return res
  }
}

window.List = List





