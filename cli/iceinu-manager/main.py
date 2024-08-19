# 一个使用Python的Prompt Toolkit编写的交互式命令行（REPL）工具
# 用于和运行中的Iceinu Bot的远程管理API进行交互
from __future__ import unicode_literals
from prompt_toolkit import prompt
from prompt_toolkit.history import InMemoryHistory


def main():
    history = InMemoryHistory()
    print("Iceinu Manager v1.0.0")
    while True:
        text = prompt("> ", history=history)
        print('You entered:', text)


if __name__ == '__main__':
    main()
