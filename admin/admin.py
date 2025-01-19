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

    def del_service(self, name: str) -> None:
        if name == "":
            raise Exception("Service name cannot be empty")

        self.DELETE("/v1/admin/service/del?name=%s" % quote_plus(name))

    def add_project(self, project: Dict[str, str]):
        if "name" not in project or project["name"] == "":
            raise Exception('Project structure is missing required "name" field')

        if "desc" not in project:
            raise Exception('Project structure is missing required "desc" field')

        if not self._check_multilang_field(project["desc"]):
            raise Exception(
                'Project structure field "desc" needs at least '
                + "one supported language entry"
            )

        self.PUT("/v1/admin/project/add", project)

    def del_project(self, name: str) -> None:
        if name == "":
            raise Exception("Project name cannot be empty")

        self.DELETE("/v1/admin/project/del?name=%s" % quote_plus(name))

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

    def del_news(self, id: str) -> None:
        if id == "":
            raise Exception("News ID cannot be empty")

        self.DELETE("/v1/admin/news/del?id=%s" % quote_plus(id))

    def logs(self) -> List[Dict[str, Any]]:
        return self.GET("/v1/admin/logs")


class AdminScript:
    def __init__(self):
        self.log: Log = Log()
        self.api: AdminAPI = None
        self.commands = {
            "add_service": self.add_service,
            "del_service": self.del_service,
            "add_project": self.add_project,
            "del_project": self.del_project,
            "add_news": self.add_news,
            "del_news": self.del_news,
            "check_services": self.check_services,
            "logs": self.get_logs,
        }
        self.api_url_env = "API_URL"

    def __format_time(self, ts: int) -> str:
        return datetime.fromtimestamp(ts, UTC).strftime("%H:%M:%S %d/%m/%Y")

    def __load_json_file(self, file: str) -> Dict[str, Any]:
        with open(file, "r") as f:
            data = loads(f.read())
            return data

    def __dump_json_file(self, data: Dict[str, Any], file: str) -> None:
        with open(file, "w") as f:
            data = dumps(data, indent=2)
            f.write(data)

    def run(self) -> bool:
        if len(argv) < 2 or len(argv) > 3:
            self.log.error("Usage: %s [command] <file>" % argv[0])
            self.log.info("Here is a list of available commands:")

            for command in self.commands.keys():
                print("\t%s" % command)

            return False

        url = getenv(self.api_url_env)
        valid_cmd = False

        if url is None:
            self.log.error(
                "Please specify the API URL using %s environment variable"
                % self.api_url_env
            )
            return False

        for cmd in self.commands:
            if argv[1] == cmd:
                valid_cmd = True
                break

        if not valid_cmd:
            self.log.error(
                "Invalid command, run the script with no commands to list the available commands"
            )
            return False

        try:
            password = self.log.password("Please enter the admin password")
            self.api = AdminAPI(url, password)

            if len(argv) == 2:
                self.handle_command(argv[1])

            elif len(argv) == 3:
                self.handle_command(argv[1], argv[2])

        except KeyboardInterrupt:
            self.log.error("Command cancelled")
            return False

        except Exception as e:
            self.log.error("Command failed: %s" % e)
            return False

    # service commands
    def add_service(self, data: Dict[str, Any] = None) -> None:
        if data is None:
            data: Dict[str, str] = {}
            data["desc"] = {}

            data["name"] = self.log.input("Serivce name")

            for lang in self.api.languages:
                data["desc"][lang] = self.log.input("Serivce desc (%s)" % lang)

            data["check_url"] = self.log.input("Serivce status check URL")
            data["clear"] = self.log.input("Serivce clearnet URL")
            data["onion"] = self.log.input("Serivce onion URL")
            data["i2p"] = self.log.input("Serivce I2P URL")

        self.api.add_service(data)
        self.log.info("Service has been added")

    def del_service(self, data: Dict[str, Any] = None) -> None:
        if data is None:
            data: Dict[str, str] = {}
            data["name"] = self.log.input("Service name")

        self.api.del_service(data["name"])
        self.log.info("Service has been deleted")

    # project commands
    def add_project(self, data: Dict[str, Any] = None) -> None:
        if data is None:
            data: Dict[str, str] = {}
            data["desc"] = {}

            data["name"] = self.log.input("Project name")

            for lang in self.api.languages:
                data["desc"][lang] = self.log.input("Project desc (%s)" % lang)

            data["url"] = self.log.input("Project URL")
            data["license"] = self.log.input("Project license")

        self.api.add_project(data)
        self.log.info("Project has been added")

    def del_project(self, data: Dict[str, Any] = None) -> None:
        if data is None:
            data: Dict[str, str] = {}
            data["name"] = self.log.input("Project name")

        self.api.del_project(data["name"])
        self.log.info("Project has been deleted")

    # news command
    def add_news(self, data: Dict[str, Any] = None) -> None:
        if data is None:
            news: Dict[str, str] = {}
            news["title"] = {}
            news["content"] = {}

            data["id"] = self.log.input("News ID")

            for lang in self.api.languages:
                data["title"][lang] = self.log.input("News title (%s)" % lang)

            data["author"] = self.log.input("News author")

            for lang in self.api.languages:
                data["content"][lang] = self.log.input("News content (%s)" % lang)

        self.api.add_news(data)
        self.log.info("News has been added")

    def del_news(self, data: Dict[str, Any] = None) -> None:
        if data is None:
            data: Dict[str, str] = {}
            data["id"] = self.log.input("News ID")

        self.api.del_project(data["id"])
        self.log.info("News has been deleted")

    def check_services(self, data: Dict[str, Any] = None) -> None:
        self.api.check_services()
        self.log.info("Requested status check for all the services")

    def get_logs(self, data: Dict[str, Any] = None) -> None:
        logs = self.api.logs()

        if logs["result"] is None or len(logs["result"]) == 0:
            return self.log.info("No available logs")

        for log in logs["result"]:
            self.log.info(
                "Time: %s | Action: %s"
                % (self.__format_time(log["time"]), log["action"])
            )

    def handle_command(self, cmd: str, file: str = None) -> bool:
        for command in self.commands.keys():
            if command != cmd:
                continue

            data = None

            try:
                if file != "" and file is not None:
                    data = self.__load_json_file(file)

                self.commands[cmd](data)
                return True

            except Exception as e:
                self.log.error("Command failed: %s" % e)
                return False

        self.log.error("Invalid command: %s", cmd)
        return False


if __name__ == "__main__":
    script = AdminScript()
    exit(script.run() if 1 else 0)
