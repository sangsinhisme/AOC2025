import json
import sys
from pulp import *

data = json.loads(sys.stdin.read())
buttons = [tuple(b) for b in data["buttons"]]
target = tuple(data["target"])

prob = LpProblem("ButtonPress", LpMinimize)
presses = [LpVariable(f"button_{i}", lowBound=0, cat='Integer') for i in range(len(buttons))]
prob += lpSum(presses)

for idx in range(len(target)):
    affecting = [presses[i] for i in range(len(buttons)) if idx in buttons[i]]
    prob += lpSum(affecting) == target[idx]

prob.solve(PULP_CBC_CMD(msg=0))

if LpStatus[prob.status] == 'Optimal':
    print(int(value(prob.objective)))
else:
    print(-1)