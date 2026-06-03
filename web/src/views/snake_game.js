// 获取画布和上下文
const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');

// 定义方块大小
const blockSize = 20;

// 定义蛇的初始位置和方向
let snake = [
    { x: 10 * blockSize, y: 10 * blockSize }
];
let direction = 'right';

// 定义食物的初始位置
let food = {
    x: Math.floor(Math.random() * (canvas.width / blockSize)) * blockSize,
    y: Math.floor(Math.random() * (canvas.height / blockSize)) * blockSize
};

// 定义分数
let score = 0;

// 游戏主循环
function gameLoop() {
    // 清除画布
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // 绘制蛇
    drawSnake();

    // 绘制食物
    drawFood();

    // 移动蛇
    moveSnake();

    // 检查碰撞
    checkCollision();

    // 检查是否吃到食物
    checkEatFood();

    // 绘制分数
    drawScore();

    // 循环调用游戏主循环
    setTimeout(gameLoop, 100);
}

// 绘制蛇
function drawSnake() {
    snake.forEach(segment => {
        ctx.fillStyle = 'green';
        ctx.fillRect(segment.x, segment.y, blockSize, blockSize);
    });
}

// 绘制食物
function drawFood() {
    ctx.fillStyle = 'red';
    ctx.fillRect(food.x, food.y, blockSize, blockSize);
}

// 移动蛇
function moveSnake() {
    let head = { ...snake[0] };
    switch (direction) {
        case 'right':
            head.x += blockSize;
            break;
        case 'left':
            head.x -= blockSize;
            break;
        case 'up':
            head.y -= blockSize;
            break;
        case 'down':
            head.y += blockSize;
            break;
    }
    snake.unshift(head);
    snake.pop();
}

// 检查碰撞
function checkCollision() {
    let head = snake[0];
    // 检查是否撞到边界
    if (head.x < 0 || head.x >= canvas.width || head.y < 0 || head.y >= canvas.height) {
        alert('游戏结束！你的分数是：' + score);
        location.reload();
    }
    // 检查是否撞到自己
    for (let i = 1; i < snake.length; i++) {
        if (head.x === snake[i].x && head.y === snake[i].y) {
            alert('游戏结束！你的分数是：' + score);
            location.reload();
        }
    }
}

// 检查是否吃到食物
function checkEatFood() {
    let head = snake[0];
    if (head.x === food.x && head.y === food.y) {
        // 增加分数
        score++;
        // 增加蛇的长度
        snake.push({});
        // 生成新的食物
        food = {
            x: Math.floor(Math.random() * (canvas.width / blockSize)) * blockSize,
            y: Math.floor(Math.random() * (canvas.height / blockSize)) * blockSize
        };
    }
}

// 绘制分数
function drawScore() {
    ctx.fillStyle = 'black';
    ctx.font = '20px Arial';
    ctx.fillText('分数: ' + score, 10, 30);
}

// 监听键盘事件
document.addEventListener('keydown', function (event) {
    switch (event.key) {
        case 'ArrowRight':
            if (direction !== 'left') {
                direction = 'right';
            }
            break;
        case 'ArrowLeft':
            if (direction !== 'right') {
                direction = 'left';
            }
            break;
        case 'ArrowUp':
            if (direction !== 'down') {
                direction = 'up';
            }
            break;
        case 'ArrowDown':
            if (direction !== 'up') {
                direction = 'down';
            }
            break;
    }
});

// 开始游戏
gameLoop();