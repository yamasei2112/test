n = int(input())

stack = []
current_page = "blank page"
stack.append(current_page)

for i in range(n):
    query = input()

    if query == "use the back button":
        if len(stack) > 1:
            stack.pop()
            current_page = stack[-1]
    elif query.startswith("go to"):
        page_name = query[6:]
        stack.append(page_name)
        current_page = page_name

    print(current_page)