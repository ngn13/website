#!/bin/python3

"""
Administration script for my website (ngn.tf)
#############################################
I really enjoy doing stuff from the terminal, 
so I wrote this simple python script that interacts 
with the API and lets me add/remove new posts/services 
from the terminal
"""

from os import remove, getenv
from getpass import getpass
import requests as req
from sys import argv

URL = "" 

def join(pth: str) -> str:
    if URL == None:
        return ""

    if URL.endswith("/"):
        return URL+pth
    return URL+"/"+pth

def get_token() -> str:
    try:
        f = open("/tmp/wa", "r")
        token = f.read()
        f.close()
        return token
    except:
        print("[-] You are not authenticated")
        exit(1)

def login() -> None:
    pwd = getpass("[>] Enter your password: ")
    res = req.get(join("admin/login")+f"?pass={pwd}").json()
    if res["error"] != "":
        print(f"[-] Error logging in: {res['error']}")
        return 

    token = res["token"]
    f = open("/tmp/wa", "w")
    f.write(token)
    f.close()

def logout() -> None:
    token = get_token()
    res = req.get(join("admin/logout"), headers={
        "Authorization": token
    }).json()
    if res["error"] != "":
        print(f"[-] Error logging out: {res['error']}")
        return

    remove("/tmp/wa")
    print("[+] Logged out")

def add_post() -> None:
    token = get_token()
    title = input("[>] Post title: ")
    author = input("[>] Post author: ")
    content_file = input("[>] Post content file: ")
    public = input("[>] Should post be public? (y/n): ")

    try:
        f = open(content_file, "r")
        content = f.read()
        f.close()
    except:
        print("[-] Content file not found")
        return

    res = req.put(join("admin/blog/add"), json={
        "title": title,
        "author": author,
        "content": content,
        "public": 1 if public == "y" else 0
    }, headers={
        "Authorization": token
    }).json()

    if res["error"] != "":
        print(f"[-] Error adding post: {res['error']}")
        return  

    print("[+] Post has been added")

def remove_post() -> None:
    token = get_token()
    id = input("[>] Post ID: ")
    res = req.delete(join("admin/blog/remove")+f"?id={id}", headers={
        "Authorization": token
    }).json()

    if res["error"] != "":
        print(f"[-] Error removing post: {res['error']}")
        return

    print("[-] Post has been removed")

def add_service() -> None:
    token = get_token()
    name = input("[>] Serivce name: ")
    desc = input("[>] Serivce desc: ")
    link = input("[>] Serivce URL: ")

    res = req.put(join("admin/service/add"), json={
        "name": name,
        "desc": desc,
        "url":  link
    }, headers={
        "Authorization": token
    }).json()

    if res["error"] != "":
        print(f"[-] Error adding service: {res['error']}")
        return

    print("[+] Service has been added")

def remove_service() -> None:
    token = get_token()
    name = input("[>] Service name: ")
    res = req.delete(join("admin/service/remove")+f"?name={name}", headers={
        "Authorization": token
    }).json()

    if res["error"] != "":
        print(f"[-] Error removing service: {res['error']}")
        return

    print("[+] Serivce has been removed")

cmds = {
    "login": login,
    "logout": logout,
    "add_post": add_post,
    "remove_post": remove_post,
    "add_service": add_service,
    "remove_service": remove_service,
}

def main():
    global URL
    URL = getenv("API")
    if URL == None or URL == "":
        print("[-] API enviroment variable not set")
        exit(1)

    if len(argv) != 2:
        print(f"[-] Usage: admin_script <command>")
        print(f"[+] Run \"admin_script help\" to get all commands")
        exit(1)

    if argv[1] == "help":
        print("Avaliable commands:")
        for k in cmds.keys():
            print(f"   {k}")
        exit()

    for k in cmds.keys():
        if k != argv[1]:
            continue 
        try:
            cmds[k]()
        except KeyboardInterrupt:
            pass
        exit()

    print("[-] Command not found")

if __name__ == "__main__":
    main()
