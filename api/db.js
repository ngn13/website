const fs = require("fs");

class Database {
  constructor() {
    this.path = "db.json";
    this.json = {};
    this.read();
  }

  write() {
    fs.writeFileSync(this.path, JSON.stringify(this.json));
  }

  read() {
    try {
      const data = fs.readFileSync(this.path, "utf8");
      this.json = JSON.parse(data);
    } catch (error) {
      return;
    }
  }

  find_all(key, check, ...args) {
    try {
      const ret = [];
      for (const d of this.json[key]) {
        if (check(d, ...args)) {
          ret.push(d);
        }
      }
      return ret;
    } catch (error) {
      return false;
    }
  }

  find(key, check, ...args) {
    try {
      for (const d of this.json[key]) {
        if (check(d, ...args)) {
          return d;
        }
      }
      return false;
    } catch (error) {
      return false;
    }
  }

  get(key) {
    const res = this.json[key]
    if(res===undefined)
      return []
    return res
  }

  push(key, data) {
    try {
      this.json[key].push(data);
    } catch (error) {
      this.json[key] = [];
      this.json[key].push(data);
    }
    this.write();
  }

  pop(key, data) {
    try {
      const indx = this.json[key].indexOf(data);
      this.json[key].splice(indx, 1);
    } catch (error) {
      return;
    }
    this.write();
  }
}

module.exports = Database;