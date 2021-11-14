import { Node, List } from './linked-list.js'

export class DoubleLinkedNode extends Node {
  pre;
  key;
  constructor (data, key) {
    super(data)
    this.pre = null
    this.key = key
  }
}

export class LRUCache extends List {
  capacity // 容量
  values // map<key, node>

  constructor(capacity) {
    super()
    this.capacity = capacity
    this.values = {}
  }

  put(key, val) {
    if (this.values[key]) {
      this.deleteNode(this.values[key])
    }

    const node = new DoubleLinkedNode(val, key)
    if (Object.keys(this.values).length === this.capacity) {
      this.deleteLast()
    }

    if (!this.end) {
      this.end = node
    }

    if (this.start) {
      node.next = this.start
      this.start.pre = node
    }

    this.start = node
    this.values[key] = node
  }

  deleteNode(node) {
    if (node.pre) {
      node.pre.next = node.next
    }

    if (this.end === node) {
      if (node.pre) {
        this.end = node.pre
      }
    }
    delete this.values[node.key]
  }

  deleteLast() {
    if (!this.end) {
      return
    }
    this.deleteNode(this.end)
  }

  get(key) {
    if (this.values[key]) {
      const node = this.values[key]
      // 删除
      if (node.pre) {
        node.pre.next = node.next
      }

      if (node.next) {
        if (node.next.pre && node.pre) {
          node.next.pre = node.pre
        }
      } else {
        if (node.pre) {
          this.end = node.pre
        }
      }

      // 放到开始
      if (this.start && this.start !== node) {
        node.next = this.start
        this.start.pre = node
        node.pre = null
        this.start = node
      }

      return this.values[key].data
    }
    return -1
  }

  display() {
    console.log(`start ${this.start.data} end ${this.end.data}`)
    this.each((item) => {
      console.log({
        data: item.data,
        pre: item.pre ? item.pre.data : null,
        next: item.next ? item.next.data : null
      })
    })
  }
}
