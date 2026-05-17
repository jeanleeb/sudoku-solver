document.addEventListener("DOMContentLoaded", () => {
  const solveBtn = document.getElementById("solve-btn");
  if (!solveBtn) return;

  solveBtn.addEventListener("click", () => {
    solveBtn.disabled = true;
    solveBtn.textContent = "Solving...";
    const inputs = document.querySelectorAll("input[name^='cell-']");
    inputs.forEach((input) => (input.disabled = true));

    const eventSource = new EventSource("/solve/stream");

    eventSource.addEventListener("step", (event) => {
      const step = JSON.parse(event.data);
      const cellId = `cell-${step.row}-${step.col}`;
      const cell = document.querySelector(`[name="${cellId}"]`);

      if (step.action == "backtrack") {
        if (cell) cell.value = "";
      } else if (step.action == "place") {
        if (cell) cell.value = step.value;
      }
    });

    eventSource.addEventListener("done", () => {
      eventSource.close();
      inputs.forEach((input) => (input.disabled = false));
      solveBtn.disabled = false;
      solveBtn.textContent = "Solve";
    });

    eventSource.addEventListener("error", () => {
      eventSource.close();
      inputs.forEach((input) => (input.disabled = false));
      solveBtn.disabled = false;
      solveBtn.textContent = "Solve";
    });
  });
});
