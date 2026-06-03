<template>
  <div style="width: 100%; height: 260px; position: reactive">
    <!-- 显示小地址小框框 -->

    <div class="ipt flex-row width100">
      <input type="text" id="suggestId" :placeholder="translang('搜索地名')" class="flex-item" />
      <div class="ml-2">
        <n-button type="primary" @click="onsubdata()" size="small">{{
          translang('提交定位')
        }}</n-button>
      </div>
    </div>
    <!-- 地图 -->
    <img
      style="width: 100%; height: 260px; object-fit: cover"
      class="mt-3 mb-5"
      src="@/assets/images/nomap.png"
      v-if="!latitude"
    />
    <div v-else id="map" :style="{ width: '100%', height: heightmap + 'px' }"> </div>
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
  import { translang } from '@/utils/smjcomm';
  import { onMounted, ref } from 'vue';

  // 传入的中心点 刷新页面加载的点
  const props = defineProps({
    latitude: {
      type: String,
      default: () => {
        return '';
      },
    },
    longitude: {
      type: String,
      default: () => {
        return '';
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
  // 显示地址
  const address = ref('');
  let adddata = ref({});
  let myValuetext = '';
  // 引入的地图的api
  const BMap = (window as any).BMap;

  function onsubdata() {
    adddata.myValuetext = myValuetext;

    emits('addressData', adddata);
  }
  const map = ref(null);

  onMounted(() => {
    if (!props.latitude) {
      return;
    }
    // 创建地图实例
    map.value = new BMap.Map('map');
    // 这个创建地理服务的 下面是把定位转换为详细文字地址
    var geoc = new BMap.Geocoder();
    //建立一个自动完成的对象 主要关键字查询
    var ac = new BMap.Autocomplete({
      // suggestId 是输入框id
      input: 'suggestId',
      //  这个是地图实例
      location: map,
    });
    // 下拉列表里的内容确认发生的事件 ()
    ac.addEventListener('onconfirm', function (e: any) {
      // 把城市啥的拼接起来
      const myValue =
        e.item.value.province +
        e.item.value.city +
        e.item.value.district +
        e.item.value.street +
        e.item.value.business;
      // 搜索
      // 搜索结束执行的函数
      const mySearchFun = () => {
        // 传入定位函数的经纬度
        getAddOverlay(local.getResults().getPoi(0).point, true);
      };
      // 创建一个搜索的实例
      var local = new BMap.LocalSearch(map, {
        //搜索成功后的回调
        onSearchComplete: mySearchFun,
      });
      // 传入搜索位置的关键字
      myValuetext = myValue;
      local.search(myValue);
    });

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

      centerAndZoom && map.value.centerAndZoom(point, 18); // 中心点位 15是级别
      map.value.addEventListener('zoomend', () => {
        console.log('当前地图缩放级别：', map.value.getZoom());
      });
      // 把定位转换为详细文字地址
      geoc.getLocation(point, (rs: any) => {
        // 显示到页面上
        address.value = rs.address;
      });
      // 把位置传出
      adddata = point;
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
    width: 420px;
    position: absolute;
    z-index: 300;
    background: #fff;
    padding: 5px;
    top: 15px;
    left: 237px;
    box-shadow: 4px 3px 5px #999999a8;
    height: 38px;

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
