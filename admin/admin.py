#!/bin/python3

"""

website/admin | Administration script for my personal website
written by ngn (https://ngn.tf) (2025)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

"""

from urllib.parse import quote_plus
from typing import Dict, List, Any
from datetime import datetime, UTC
from colorama import Fore, Style
from json import dumps, loads
from getpass import getpass
import requests as req
from os import getenv
from sys import argv

API_URL_ENV = "API_URL"


# logger used by the script
class Log:
    def __init__(self) -> None:
        self.reset = Fore.RESET + Style.RESET_ALL

    def info(self, m: str) -> None:
        print(Fore.BLUE + Style.BRIGHT + "[*]" + self.reset + " " + m)

    def error(self, m: str) -> None:
        print(Fore.RED + Style.BRIGHT + "[-]" + self.reset + " " + m)

    def input(self, m: str) -> str:
        return input(Fore.CYAN + Style.BRIGHT + "[?]" + self.reset + " " + m + ": ")

    def password(self, m: str) -> str:
        return getpass(Fore.CYAN + Style.BRIGHT + "[?]" + self.reset + " " + m + ": ")


# API interface for the admin endpoints
class AdminAPI:
    def __init__(self, url: str, password: str) -> None:
        self.languages: List[str] = [
            "en",
            "tr",
        ]  # languages supported by multilang fields
        self.password = password
        self.api_url = url

    def _title_to_id(self, title: str) -> str:
        return title.lower().replace(" ", "_")

    def _check_multilang_field(self, ml: Dict[str, str]) -> bool:
        for lang in self.languages:
            if lang in ml and ml[lang] != "":
                return True
        return False

    def _api_url_join(self, path: str) -> str:
        api_has_slash = self.api_url.endswith("/")
        path_has_slash = path.startswith("/")

        if api_has_slash or path_has_slash:
            return self.api_url + path
        elif api_has_slash and path_has_slash:
            return self.api_url + path[1:]
        else:
            return self.api_url + "/" + path

    def _to_json(self, res: req.Response) -> dict:
        if res.status_code == 403:
            raise Exception("Authentication failed")

        json = res.json()

        if json["error"] != "":
            raise Exception("API error: %s" % json["error"])

        return json

    def PUT(self, url: str, data: dict) -> req.Response:
        return self._to_json(
            req.put(
                self._api_url_join(url),
                json=data,
                headers={"Authorization": self.password},
            )
        )

    def DELETE(self, url: str) -> req.Response:
        return self._to_json(
            req.delete(
                self._api_url_join(url), headers={"Authorization": self.password}
            )
        )

    def GET(self, url: str) -> req.Response:
        return self._to_json(
            req.get(self._api_url_join(url), headers={"Authorization": self.password})
        )

    def add_service(self, service: Dict[str, str]):
        if "name" not in service or service["name"] == "":
            raise Exception('Service structure is missing required "name" field')

        if "desc" not in service:
            raise Exception('Service structure is missing required "desc" field')

        if (
            ("clear" not in service or service["clear"] == "")
            and ("onion" not in service or service["onion"] == "")
            and ("i2p" not in service or service["i2p"] == "")
        ):
            raise Exception(
                'Service structure is missing "clear", "onion" '
                + 'and "i2p" field, at least one needed'
            )

        if not self._check_multilang_field(service["desc"]):
            raise Exception(
                'Service structure field "desc" needs at least '
                + "one supported language entry"
            )

        self.PUT("/v1/admin/service/add", service)

    def del_service(self, service: str) -> None:
        if service == "":
            raise Exception("Service name cannot be empty")

        self.DELETE("/v1/admin/service/del?name=%s" % quote_plus(service))

    def check_services(self) -> None:
        self.GET("/v1/admin/service/check")

    def add_news(self, news: Dict[str, str]):
        if "id" not in news or news["id"] == "":
            raise Exception('News structure is missing required "id" field')

        if "author" not in news or news["author"] == "":
            raise Exception('News structure is missing required "author" field')

        if "title" not in news:
            raise Exception('News structure is missing required "title" field')

        if "content" not in news:
            raise Exception('News structure is missing required "content" field')

        if not self._check_multilang_field(news["title"]):
            raise Exception(
                'News structure field "title" needs at least '
                + "one supported language entry"
            )

        if not self._check_multilang_field(news["content"]):
            raise Exception(
                'News structure field "content" needs at least '
                + "one supported language entry"
            )

        self.PUT("/v1/admin/news/add", news)

    def del_news(self, news: str) -> None:
        if news == "":
            raise Exception("News ID cannot be empty")

        self.DELETE("/v1/admin/news/del?id=%s" % quote_plus(news))

    def logs(self) -> List[Dict[str, Any]]:
        return self.GET("/v1/admin/logs")


# local helper functions used by the script
def __format_time(ts: int) -> str:
    return datetime.fromtimestamp(ts, UTC).strftime("%H:%M:%S %d/%m/%Y")


def __load_json_file(file: str) -> Dict[str, Any]:
    with open(file, "r") as f:
        data = loads(f.read())
        return data


def __dump_json_file(data: Dict[str, Any], file: str) -> None:
    with open(file, "w") as f:
        data = dumps(data, indent=2)
        f.write(data)


# command handlers
def __handle_command(log: Log, api: AdminAPI, cmd: str) -> None:
    match cmd:
        case "add_service":
            data: Dict[str, str] = {}
            data["desc"] = {}

            data["name"] = log.input("Serivce name")
            for lang in api.languages:
                data["desc"][lang] = log.input("Serivce desc (%s)" % lang)
            data["check_url"] = log.input("Serivce status check URL")
            data["clear"] = log.input("Serivce clearnet URL")
            data["onion"] = log.input("Serivce onion URL")
            data["i2p"] = log.input("Serivce I2P URL")

            api.add_service(data)
            log.info("Service has been added")

        case "del_service":
            api.del_service(log.input("Serivce name"))
            log.info("Service has been deleted")

        case "check_services":
            api.check_services()
            log.info("Requested status check for all the services")

        case "add_news":
            news: Dict[str, str] = {}
            news["title"] = {}
            news["content"] = {}

            data["id"] = log.input("News ID")
            for lang in api.languages:
                data["title"][lang] = log.input("News title (%s)" % lang)
            data["author"] = log.input("News author")
            for lang in api.languages:
                data["content"][lang] = log.input("News content (%s)" % lang)

            api.add_news(data)
            log.info("News has been added")

        case "del_news":
            api.del_news(log.input("News ID"))
            log.info("News has been deleted")

        case "logs":
            logs = api.logs()

            if logs["result"] is None or len(logs["result"]) == 0:
                return log.info("No available logs")

            for log in logs["result"]:
                log.info(
                    "Time: %s | Action: %s"
                    % (__format_time(log["time"]), log["action"])
                )


def __handle_command_with_file(log: Log, api: AdminAPI, cmd: str, file: str) -> None:
    match cmd:
        case "add_service":
            data = __load_json_file(file)
            api.add_service(data)
            log.info("Service has been added")

        case "del_service":
            data = __load_json_file(file)
            api.del_service(data["name"])
            log.info("Service has been deleted")

        case "check_services":
            api.check_services()
            log.info("Requested status check for all the services")

        case "add_news":
            data = __load_json_file(file)
            api.add_news(data)
            log.info("News has been added")

        case "del_news":
            data = __load_json_file(file)
            api.del_news(data["id"])
            log.info("News has been deleted")

        case "logs":
            logs = api.logs()

            if logs["result"] is None or len(logs["result"]) == 0:
                return log.info("No available logs")

            __dump_json_file(logs["result"], file)
            log.info("Logs has been saved")


if __name__ == "__main__":
    log = Log()

    if len(argv) < 2 or len(argv) > 3:
        log.error("Usage: %s [command] <file>" % argv[0])
        log.info("Here is a list of available commands:")
        print("\tadd_service")
        print("\tdel_service")
        print("\tcheck_services")
        print("\tadd_news")
        print("\tdel_news")
        print("\tlogs")
        exit(1)

    url = getenv(API_URL_ENV)

    if url is None:
        log.error(
            "Please specify the API URL using %s environment variable" % API_URL_ENV
        )
        exit(1)

    try:
        password = log.password("Please enter the admin password")
        api = AdminAPI(url, password)

        if len(argv) == 2:
            __handle_command(log, api, argv[1])
        elif len(argv) == 3:
            __handle_command_with_file(log, api, argv[1], argv[2])

    except KeyboardInterrupt:
        print()
        log.error("Command cancelled")
        exit(1)

    except Exception as e:
        log.error("Command failed: %s" % e)
        exit(1)
