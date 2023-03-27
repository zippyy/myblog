<!DOCTYPE html>
<!-- Created By MultiWebPress - www.multiwebpress.com -->
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Poll UI Design | MultiWebPress</title>
  <link rel="stylesheet" href="css/style.css">
</head>
<body>
  <div class="wrapper">
    <header>What Design tool do you use the most? <br></header>
    <div class="poll-area">
      <input type="checkbox" name="poll" id="opt-1">
      <input type="checkbox" name="poll" id="opt-2">
      <input type="checkbox" name="poll" id="opt-3">
      <input type="checkbox" name="poll" id="opt-4">
      <label for="opt-1" class="opt-1">
        <div class="row">
          <div class="column">
            <span class="circle"></span>
            <span class="text">Photoshop</span>
          </div>
          <span class="percent">55%</span>
        </div>
        <div class="progress" id="pstyle1" style='--w:55;'></div>
      </label>
      <label for="opt-2" class="opt-2">
        <div class="row">
          <div class="column">
            <span class="circle"></span>
            <span class="text">Sketch</span>
          </div>
          <span class="percent">20%</span>
        </div>
        <div class="progress" id="pstyle2" style='--w:80;'></div>
      </label>
      <label for="opt-3" class="opt-3">
        <div class="row">
          <div class="column">
            <span class="circle"></span>
            <span class="text">Adobe XD</span>
          </div>
          <span class="percent">20%</span>
        </div>
        <div class="progress" id="pstyle3" style='--w:20;'></div>
      </label>
      <label for="opt-4" class="opt-4">
        <div class="row">
          <div class="column">
            <span class="circle"></span>
            <span class="text">Figma</span>
          </div>
          <span class="percent">96%</span>
        </div>
        <div class="progress" id="pstyle4" style='--w:96;'></div>
      </label>
    </div>
  </div>
  <script src="javascript/script.js"></script>
</body>
</html>