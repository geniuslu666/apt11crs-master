<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">首页活动详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <n-grid y-gap="25" :cols="3">
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">标题</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.title }}
                  </div>
                </div>
              </n-gi>
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">作者</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.author }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">语言</div>
                  <div class="level-detail-div-item-content1" :class="getOptionTag(options.language, formValue?.language)">{{ getOptionLabel(options.language, formValue?.language) }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">浏览量</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.views }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">点赞量</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.thumbNum }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">排序</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.sort }}
                  </div>
                </div>
              </n-gi>
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">内容</div>
                  <div class="level-detail-div-item-content">
                    <div v-html="formValue.content"></div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">券链接区域标题</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.couponLinkTitle }}
                  </div>
                </div>
              </n-gi>
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">券链接区域副标题</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.couponLinkSubTitle }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">推荐链接区域标题</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.recommendLinkTitle }}
                  </div>
                </div>
              </n-gi>
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">推荐链接区域副标题</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.recommendLinkSubTitle }}
                  </div>
                </div>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('foodIndex') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">餐厅首页推荐标题</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.foodIndexTitle }}
                  </div>
                </div>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('spaIndex') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">按摩首页推荐标题</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.spaIndexTitle }}
                  </div>
                </div>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('carIndex') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">接送机首页推荐标题</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.carIndexTitle }}
                  </div>
                </div>
              </n-gi>
              <n-gi span="3" v-if="recommendLinkTypeArr.indexOf('coupon') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">推荐优惠券</div>
                  <div class="level-detail-div-item-content">
                    <n-table style="margin-bottom: 25px">
                      <thead>
                      <tr>
                        <th>优惠券名称</th>
                        <th>优惠券类型</th>
                        <th>优惠券金额/折扣</th>
                        <th>满多少元使用</th>
                        <th>适用场景</th>
                        <th>有效期</th>
                        <th>每人领取数量</th>
                      </tr>
                      </thead>
                      <tbody>
                      <tr v-for="(item, index) of couponArr" :key="index">
                        <td>{{ item.couponName }}</td>
                        <td>
                          <template v-if="item.type == 'reward'">
                            <span style="color: green">满减</span>
                          </template>
                          <template v-if="item.type == 'discount'">
                            <span style="color: blue">折扣</span>
                          </template>
                        </td>
                        <td>
                          <template v-if="item.type == 'reward'">
                            {{ item.money }}JPY
                          </template>
                          <template v-if="item.type == 'discount'">
                            {{ item.discount }}折
                          </template>
                        </td>
                        <td>{{ item.atLeast }}</td>
                        <td>
                          <template v-if="item.scene == 1">
                            <span style="color: #26A763">民宿</span>
                          </template>
                          <template v-if="item.scene == 2">
                            <span style="color: blue">餐饮</span>
                          </template>
                          <template v-if="item.scene == 3">
                            <span style="color: #EFA020">按摩</span>
                          </template>
                          <template v-if="item.scene == 4">
                            <span style="color: #F56C6C">接送机/包车</span>
                          </template>
                          <template v-if="item.scene == 5">
                            <span style="color: #F56C6C">储物柜</span>
                          </template>
                        </td>
                        <td>
                          <template v-if="item.validityType == 1">
                            {{ item.endUseTime }}
                          </template>
                          <template v-else-if="item.validityType == 2">
                            领取后{{ item.fixedTerm }}天有效
                          </template>
                          <template v-else>
                            长期有效
                          </template>
                        </td>
                        <td>
                          {{ item.num }}
                        </td>
                      </tr>
                      </tbody>
                    </n-table>
                  </div>
                </div>
              </n-gi>

              <n-gi span="3" v-if="recommendLinkTypeArr.indexOf('thCoupon') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">礼品券</div>
                  <div class="level-detail-div-item-content">
                    <n-table style="margin-bottom: 25px">
                      <thead>
                      <tr>
                        <th>礼品券名称</th>
                        <th>适用商户</th>
                        <th>每人每天领取数</th>
                        <th>每人领取数量</th>
                      </tr>
                      </thead>
                      <tbody>
                      <tr v-for="(item, index) of thCouponArr" :key="index">
                        <td>{{ item.couponName }}</td>
                        <td>
                          <n-ellipsis expand-trigger="click" :line-clamp="1" :tooltip="false" style="width: 200px">
                            <div v-for="(item1, index1) of item.mchList" :key="index1">{{ item1.mchInfo.name }}</div>
                          </n-ellipsis>
                        </td>
                        <td>
                          {{ item.perNum }}
                        </td>
                        <td>
                          {{ item.num }}
                        </td>
                      </tr>
                      </tbody>
                    </n-table>
                  </div>
                </div>
              </n-gi>
              <n-gi span="3" v-if="recommendLinkTypeArr.indexOf('hotelDetail') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">民宿</div>
                  <div class="level-detail-div-item-content">
                    <n-table style="margin-bottom: 25px">
                      <thead>
                      <tr>
                        <th>民宿名称</th>
                      </tr>
                      </thead>
                      <tbody>
                      <tr v-for="(item, index) of hotelArr" :key="index">
                        <td>{{ item.name }}</td>
                      </tr>
                      </tbody>
                    </n-table>
                  </div>
                </div>
              </n-gi>
              <n-gi span="3" v-if="recommendLinkTypeArr.indexOf('foodDetail') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">餐厅</div>
                  <div class="level-detail-div-item-content">
                    <n-table style="margin-bottom: 25px">
                      <thead>
                      <tr>
                        <th>餐厅名称</th>
                        <th>营业类型</th>
                        <th>地址</th>
                      </tr>
                      </thead>
                      <tbody>
                      <tr v-for="(item, index) of restaurantArr" :key="index">
                        <td>{{ item.name }}</td>
                        <td>{{ item.cooperateTypeId > 0 ? item.cooperateTypeDetail.typeName : '--' }}</td>
                        <td>{{ item.detailAddress }}</td>
                      </tr>
                      </tbody>
                    </n-table>
                  </div>
                </div>
              </n-gi>
              <n-gi span="3" v-if="recommendLinkTypeArr.indexOf('spaDetail') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">按摩服务</div>
                  <div class="level-detail-div-item-content">
                    <n-table style="margin-bottom: 25px">
                      <thead>
                      <tr>
                        <th>服务名称</th>
                        <th>所属服务商</th>
                      </tr>
                      </thead>
                      <tbody>
                      <tr v-for="(item, index) of spaServiceArr" :key="index">
                        <td>{{ item.name }}</td>
                        <td>{{ item.ispId > 0 ? item.spaIspDetail.name : '--' }}</td>
                      </tr>
                      </tbody>
                    </n-table>
                  </div>
                </div>
              </n-gi>
              <n-gi span="3" v-if="recommendLinkTypeArr.indexOf('outLink') !== -1">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">外部链接</div>
                  <div class="level-detail-div-item-content">
                    <n-table style="margin-bottom: 25px">
                      <thead>
                      <tr>
                        <th>外链标题</th>
                        <th>外链链接</th>
                        <th>按钮文字</th>
                        <th>跳转方式</th>
                      </tr>
                      </thead>
                      <tbody>
                      <tr v-for="(item, index) of outLinkArr" :key="index">
                        <td>{{ item.title }}</td>
                        <td>{{ item.link }}</td>
                        <td>{{ item.buttonTxt }}</td>
                        <td>{{ item.openType == 1 ? '内部webview' : '外部浏览器' }}</td>
                      </tr>
                      </tbody>
                    </n-table>
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { View } from '@/api/homepageArticles';
import { State, newState, options } from './model';
import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {ThCouponMchInter} from "/#/thCoupon";
import {All as ThCouponAll } from "@/api/thCoupon";
import {All} from "@/api/pmsCouponType";
import {All as RestaurantAll } from "@/api/foodRestaurant";
import {All as SpaServiceAll } from "@/api/spaService";
import {All as HotelAll } from "@/api/pmsProperty";

const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(800);
});
const recommendLinkTypeArr = ref<string[]>([])
// ---礼品券---
interface SelectThCouponInter {
  id: number;
  couponName: string;
  mchList: ThCouponMchInter[];
  num: number;
  perNum: number;
}

interface ThCouponIdsNumInter {
  couponId: number;
  num: number;
  perNum: number;
}
const thCouponIdsNumArr = ref<ThCouponIdsNumInter[]>([]);
const thCouponArr = ref<SelectThCouponInter[]>([])
// ---礼品券-END---

// ---优惠券---
interface CouponIdsNumInter {
  couponId: number;
  num: number;
}
interface SelectCouponInter {
  id: number;
  couponName: string;
  type: string;
  money: number;
  discount: number;
  atLeast: number;
  scene: number;
  validityType: number;
  endUseTime: string;
  fixedTerm: number;
  num: number;
}
const couponIdsNumArr = ref<CouponIdsNumInter[]>([]);
const couponArr = ref<SelectCouponInter[]>([])
// ---优惠券-END---

// ---餐厅---
interface SelectRestaurantInter {
  id: number;
  name: string;
  cooperateTypeId: number;
  cooperateTypeDetail: {
    typeName: string
  };
  detailAddress: string;
  phone: string;
}
const restaurantIdsArr = ref([]);
const restaurantArr = ref<SelectRestaurantInter[]>([])
// ---餐厅-END---

// ---按摩---
interface SelectSpaServiceInter {
  id: number;
  name: string;
  ispId: number;
  spaIspDetail: {
    name: string
  };
}
const spaServiceIdsArr = ref([]);
const spaServiceArr = ref<SelectSpaServiceInter[]>([])
// ---按摩-END---

// ---酒店---
interface SelectHotelInter {
  id: number;
  name: string;
  close: number;
}
const hotelIdsArr = ref([]);
const hotelArr = ref<SelectHotelInter[]>([])
// ---酒店-END---

// ---外链---
interface SelectOutLinkInter {
  title: string;
  buttonTxt: string;
  link: string;
  openType: number;
}
const outLinkArr = ref<SelectOutLinkInter[]>([])
// ---外链-END---

async function openModal(state: State) {
  showModal.value = true;
  loading.value = true;
  await getInfo(state.id)
  loading.value = false;
}

async function getInfo(id){
  const res = await View({ id: id, isLanguage: true });
  formValue.value = res;

  recommendLinkTypeArr.value = res.recommendLinkType.split(',')

  console.log(' recommendLinkTypeArr.value',  recommendLinkTypeArr.value)

  thCouponIdsNumArr.value = []
  thCouponArr.value = []

  couponIdsNumArr.value = []
  couponArr.value = []

  restaurantIdsArr.value = []
  restaurantArr.value = []

  spaServiceIdsArr.value = []
  spaServiceArr.value = []

  hotelIdsArr.value = []
  hotelArr.value = []

  if(res.outLinks){
    outLinkArr.value = res.outLinks
  }

  if(res.couponList){
    res.couponList.forEach((item) => {
      if(item.couponType == 'thCoupon'){
        thCouponIdsNumArr.value.push({couponId: item.couponId, num: item.availableQuantity, perNum: item.perDayAvailable})
      }else{
        couponIdsNumArr.value.push({couponId: item.couponId, num: item.availableQuantity})
      }
    })
  }
  if(thCouponIdsNumArr.value.length > 0){
    await getThCouponList()
    if(thCouponArr.value.length > 0){
      thCouponArr.value.forEach(m => {
        m.num = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.num || 1;
        m.perNum = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.perNum || 0
      })
    }
  }
  if(couponIdsNumArr.value.length > 0){
    await getCouponList()
    if(couponArr.value.length > 0){
      couponArr.value.forEach(m => {
        m.num = couponIdsNumArr.value.find(n => n.couponId == m.id)?.num || 1
      })
    }
  }

  if(res.restaurantIds){
    restaurantIdsArr.value = res.restaurantIds.split(',').map((item) => {
      return Number(item)
    })
  }else{
    restaurantIdsArr.value = []
  }
  if(restaurantIdsArr.value.length > 0){
    await getRestaurantList()
  }
  // 按摩服务
  if(res.spaServiceIds){
    spaServiceIdsArr.value = res.spaServiceIds.split(',').map((item) => {
      return Number(item)
    })
  }else{
    spaServiceIdsArr.value = []
  }
  if(spaServiceIdsArr.value.length > 0){
    await getSpaServiceList()
  }
  // 酒店
  if(res.hotelIds){
    hotelIdsArr.value = res.hotelIds.split(',').map((item) => {
      return Number(item)
    })
  }else{
    hotelIdsArr.value = []
  }
  if(hotelIdsArr.value.length > 0){
    await getHotelList()
  }
}
// 获取礼品券列表
async function getThCouponList(){
  const res = await ThCouponAll({
    couponIds: thCouponIdsNumArr.value.map(m => m.couponId).join(',')
  });
  thCouponArr.value = res.list
  console.log('thCouponArr',thCouponArr.value)
}
// 获取优惠券列表
async function getCouponList(){
  const res = await All({
    couponIds: couponIdsNumArr.value.map(m => m.couponId).join(',')
  });
  couponArr.value = res.list
  console.log('couponArr',couponArr.value)
}
// 获取餐厅列表
async function getRestaurantList(){
  const res = await RestaurantAll({
    restaurantIds: restaurantIdsArr.value.join(',')
  });
  restaurantArr.value = res.list
  console.log('restaurantArr',restaurantArr.value)
}
// 获取按摩列表
async function getSpaServiceList(){
  const res = await SpaServiceAll({
    serviceIds: spaServiceIdsArr.value.join(',')
  });
  spaServiceArr.value = res.list
  console.log('spaServiceArr',spaServiceArr.value)
}
// 获取酒店列表
async function getHotelList(){
  const res = await HotelAll({
    ids: hotelIdsArr.value.join(',')
  });
  hotelArr.value = res.list
  console.log('hotelArr',hotelArr.value)
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>
.level-detail-div{
  margin-bottom: 25px;
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 25px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      margin-bottom: 8px;
    }
    &-content{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      display: flex;
      align-items: center;
      .color{
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
      }
    }
    &-content1{
      padding: 0 6px;
      height: 24px;
      line-height: 24px;
      font-weight: 400;
      font-size: 14px;
      display: inline-block;
      &.success{
        color: #19A158;
        background: #E3F4EB;
      }
      &.info{
        color: #3F9EFF;
        background: #ECF5FF;
      }
      &.default{
        color: #919399;
        background: #F4F4F5;
      }
      &.warning{
        color: #EFA020;
        background: #FDF1DD;
      }
      &.error{
        color: #F56C6C;
        background: #FEF0F0;
      }
    }
  }
}
</style>


