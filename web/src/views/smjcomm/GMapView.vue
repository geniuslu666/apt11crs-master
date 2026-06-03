<template>
  <div>
    <img
      style="width: 100%; height: 260px; object-fit: cover"
      class="mt-3 mb-5"
      src="@/assets/images/nomap.png"
      v-if="longitude == ''"
    />
    <div v-show="longitude != ''" class="pac-card div-style" id="pac-card">
      <div id="map" class="flex-item1"></div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'addressMap',
    props: ['long', 'lat'],
    data() {
      return {
        localValue: '',
        map: '',
        marker: '',
        longitude: '',
        latitude: '',
        predictions: [], // 存储搜索框提示的预测结果
      };
    },
    methods: {
      onSelect(value) {
        console.log('onSelect', value);
      },
      selectPrediction(val) {
        let prediction = '';
        this.predictions.forEach((element) => {
          if (element.value == val) {
            prediction = element;
          }
        });

        this.localValue = prediction.description;
        this.predictions = [];
        const that = this;
        // 创建 PlacesService 对象
        const placesService = new google.maps.places.PlacesService(map);
        // 获取地点的 Place ID
        const placeId = prediction.place_id;
        // 发起 Places API 请求
        placesService.getDetails({ placeId: placeId }, function (place, status) {
          if (status === google.maps.places.PlacesServiceStatus.OK) {
            // 获取地点的经纬度坐标
            that.latitude = place.geometry.location.lat();
            that.longitude = place.geometry.location.lng();
            that.initMap();
          } else {
            alert('无法找到该地点！');
          }
        });
      },
      onsubdata() {
        this.$emit('addressData', { lat: this.latitude, lng: this.longitude });
      },
      handleInput() {
        const autocompleteService = new google.maps.places.AutocompleteService();
        autocompleteService.getPlacePredictions(
          { input: this.localValue },
          (predictions, status) => {
            if (status === google.maps.places.PlacesServiceStatus.OK) {
              this.predictions = predictions.map((m) => {
                return {
                  ...m,
                  value: m.description,
                };
              });
            } else {
              this.predictions = [];
            }
          }
        );
      },
      initMap() {
        const map = new google.maps.Map(document.getElementById('map'), {
          center: { lat: this.latitude, lng: this.longitude },
          zoom: 18,
          mapTypeControl: false,
        });
        var marker = new google.maps.Marker({
          position: new google.maps.LatLng(this.latitude, this.longitude),
        });
        marker.setMap(map);
      },
    },
    mounted() {
      let _this = this;
      axios
        .get(
          'https://maps.googleapis.com/maps/api/js?key=AIzaSyBVpL8bVshPstY8IxslE19gDWQzI00tzBs&libraries=places',
          {
            timeout: 3000, // 设置超时时间为5秒
          }
        )
        .then((response) => {
          // 请求成功处理逻辑
          const script = document.createElement('script');
          script.text = response.data;
          document.body.appendChild(script);
          _this.longitude = _this.long - 0;
          _this.latitude = _this.lat - 0;
          _this.initMap();
        })
        .catch((error) => {
          // 请求失败处理逻辑
        });
    },
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="less">
  #map {
    overflow: hidden;
    width: 100%;
    height: 260px;
    margin: 0;
  }
  .div-style {
    display: flex;
  }
  .flex-item1 {
    flex: 2;
  }
  .flex-item2 {
    flex: 1;
  }
  .autocomplete-list {
    height: 200px;
    overflow: auto;
    ul {
      li {
        margin: 10px 0;
      }
    }
  }
</style>
