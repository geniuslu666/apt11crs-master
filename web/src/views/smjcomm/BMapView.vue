<template>
  <div style="width: 100%; height: 260px">
    <!-- 地图 -->
    <div id="map" :style="{ width: '100%', height: heightmap + 'px' }"> </div>
    <div
      style="
        cursor: pointer;
        position: absolute;
        right: 10px;
        top: 10px;
        text-align: center;
        color: #053dc8;
        width: 45px;
        height: 45px;
        font-size: 40px;
        line-height: 45px;
        background: #2f2f2f24;
      "
      @click="zoomIn"
      >+
    </div>
    <div
      style="
        cursor: pointer;
        position: absolute;
        right: 10px;
        top: 60px;
        text-align: center;
        color: #053dc8;
        width: 45px;
        height: 45px;
        font-size: 40px;
        line-height: 40px;
        background: #2f2f2f24;
      "
      @click="zoomOut"
      >-
    </div>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, ref } from 'vue';

  // 传入的中心点 刷新页面加载的点
  const props = defineProps({
    latitude: {
      type: String,
      default: () => {
        return '40.083402';
      },
    },
    longitude: {
      type: String,
      default: () => {
        return '113.368374';
      },
    },
    heightmap: {
      type: Number,
      default: () => {
        return 400;
      },
    },
  });
  // 传出地址
  const emits = defineEmits(['addressData']);
  // 引入的地图的api
  const BMap = (window as any).BMap;
  const map = ref(null);

  onMounted(() => {
    // 创建地图实例
    map.value = new BMap.Map('map');
    // 这个创建地理服务的 下面是把定位转换为详细文字地址
    var geoc = new BMap.Geocoder();

    // 下面是开始定位的
    var point = new BMap.Point(props.longitude, props.latitude); // 定位
    getAddOverlay(point, true);
    map.value.enableScrollWheelZoom(true); //开启鼠标滚轮缩放
    // 当地图点击的时候发生的事件
    map.value.addEventListener('click', function (e: any) {
      // 创建标点
      getAddOverlay(new BMap.Point(e.point.lng, e.point.lat));
    });
    // 定位点的函数
    function getAddOverlay(point: any, centerAndZoom: boolean = false) {
      // 清空地图上所有的标准当然你想要多个点的话可以不清除
      map.value.clearOverlays();
      // 创建图标对象，这里的图标大小是由图片本身的尺寸决定的
      var myIcon = new BMap.Icon(
        'https://apt11-1251002327.cos.ap-tokyo.myqcloud.com/storage/public/2024-07-10/d2lmowm1wumlijfde7.png',
        new BMap.Size(30, 30)
      );

      // 使用自定义图标创建标记
      var marker = new BMap.Marker(point, { icon: myIcon });
      map.value.addOverlay(marker); // 添加到地图
      // 监听放大事件

      centerAndZoom && map.value.centerAndZoom(point, 18); // 中心点位 15是级别
      map.value.addEventListener('zoomend', () => {
        console.log('当前地图缩放级别：', map.value.getZoom());
      });
    }
  });
  const zoomIn = () => {
    if (map.value) {
      map.value.zoomIn(); // 地图放大
    }
  };

  const zoomOut = () => {
    if (map.value) {
      map.value.zoomOut(); // 地图缩小
    }
  };
</script>

<style lang="less">
  .anchorBL {
    display: none;
  }
</style>
<style lang="less" scoped>
  .ipt {
    position: relative;
    display: flex;
    justify-content: flex-end;
    width: 100%;
    min-width: 300px;
    margin-bottom: 10px;

    #suggestId {
      padding: 5px 10px;
      outline: none;
      font-size: 13px;
      font-family: monospace;
      color: #606266;
      border-radius: 4px;
      border: 1px solid #ddd;
    }

    .address {
      position: absolute;
      width: 300px;
      height: 20px;
      font-size: 14px;
      left: 0;
      bottom: 0;
    }
  }
</style>
