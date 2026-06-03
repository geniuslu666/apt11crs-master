<template>
  <div>
    <n-drawer v-model:show="showModal" :trap-focus="false" :width="dialogWidth" :z-index="99">
      <n-drawer-content closable :header-style="{
        padding: '20px',
      }" :body-content-style="{
        padding: '20px',
      }" :footer-style="{
        padding: '12px 20px',
      }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">
            {{ formValue.id > 0 ? '编辑首页活动#' + formValue.id : '添加首页活动' }}</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form ref="formRef" label-placement="top" label-width="auto" :model="formValue" :rules="rules">
            <n-grid :cols="1">
              <n-gi>
                <n-form-item label="语言" path="language">
                  <n-select placeholder="请选择语言" v-model:value="formValue.language" :options="options.language" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="活动编号" path="no" :show-require-mark="true" :show-feedback='false'>
                  <n-input v-model:value="formValue.no" placeholder="请输入编号" />
                </n-form-item>
                <div style="color:red; font-size: 12px;line-height: 34px;margin: 0 0 24px">相同的活动不同语言需要填写相同的活动编号，请谨慎填写！！！</div>
              </n-gi>
              <n-gi>
                <n-form-item label="标题" path="title">
                  <n-input v-model:value="formValue.title" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
<!--              <n-gi>
                <n-form-item label="作者" path="author">
                  <n-input v-model:value="formValue.author" placeholder="请输入作者" />
                </n-form-item>
              </n-gi>-->
              <n-gi>
                <n-form-item label="列表图" path="bannerImage" :show-require-mark="true" :show-feedback='false'>
                  <FileChooser1 v-model:value="formValue.listPic" :maxNumber="1" fileType="default" />
                </n-form-item>
                <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px">建议尺寸：670px*320px</div>
              </n-gi>
              <n-gi>
                <n-form-item label="活动开始时间" path="startTime" :show-require-mark="true">
                  <DatePicker v-model:formValue="formValue.startTime" type="date" style="width: 300px"
                    :is-date-disabled="disablePreviousDate" :disabled="formValue.id > 0" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="活动结束时间" path="endTime" :show-require-mark="true">
                  <DatePicker v-model:formValue="formValue.endTime" type="date" style="width: 300px"
                    :is-date-disabled="disablePreviousDate" :disabled="formValue.id > 0" />
                </n-form-item>
              </n-gi>
              <n-gi style="margin-bottom: 10px">
                <n-form-item label="排序" path="sort">
                  <n-input-number v-model:value="formValue.sort" clearable placeholder="请输入排序" style="width: 300px" />
                  <template #feedback>
                    <div style="font-size: 12px">越大越靠前</div>
                  </template>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button v-for="status in options.sys_normal_disable" :key="status.value"
                                    :value="status.value" :label="status.label" />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="小程序显示" path="status">
                  <n-input-group style="margin-left: 8px">
                    <n-switch v-model:value="formValue.minappStatus" :unchecked-value="2" :checked-value="1">
                      <template #checked>
                        显示
                      </template>
                      <template #unchecked>
                        不显示
                      </template>
                    </n-switch>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="活动链接方式" path="chain">
                  <n-radio-group v-model:value="formValue.chain" name="chain">
                    <n-space>
                      <n-radio value="IN">
                        内链
                      </n-radio>
                      <n-radio value="OUT">
                        外链
                      </n-radio>
                    </n-space>
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.chain == 'OUT'">
                <n-form-item label="活动跳转外链链接" path="path">
                  <n-input v-model:value="formValue.path" placeholder="请输入活动跳转外链链接" />
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.chain == 'OUT'">
                <n-form-item label="活动外链跳转方式" path="linkOpenType">
                  <n-radio-group v-model:value="formValue.linkOpenType" name="linkOpenType">
                    <n-space>
                      <n-radio :value="1">
                        内部webview
                      </n-radio>
                      <n-radio :value="2">
                        外部浏览器
                      </n-radio>
                    </n-space>
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <template v-if="formValue.chain == 'IN'">
              <n-gi>
                <n-form-item label="内容" path="content">
                  <Editor id="content" v-model:modelValue="formValue.content" style="height: 400px" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item path="rewardType" label="推荐类型" :show-require-mark="true">
                  <n-checkbox-group v-model:value="recommendLinkTypeArr" @update:value="handleUpdateValue">
                    <n-space item-style="display: flex;">
                      <n-checkbox value="coupon" label="优惠券" />
                      <n-checkbox value="thCoupon" label="礼品券" />
                      <n-checkbox value="hotelDetail" label="民宿详情" />
                      <n-checkbox value="foodIndex" label="餐厅首页" />
                      <n-checkbox value="foodDetail" label="餐厅详情" />
                      <n-checkbox value="spaIndex" label="按摩首页" />
                      <n-checkbox value="spaDetail" label="按摩详情" />
                      <n-checkbox value="carIndex" label="接送机首页" />
                      <n-checkbox value="outLink" label="外链" />
                    </n-space>
                  </n-checkbox-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="券链接区域标题" path="title" :show-require-mark="false">
                  <n-input v-model:value="formValue.couponLinkTitle" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="券链接区域副标题" path="title" :show-require-mark="false">
                  <n-input v-model:value="formValue.couponLinkSubTitle" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="推荐链接区域标题" path="title" :show-require-mark="false">
                  <n-input v-model:value="formValue.recommendLinkTitle" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="推荐链接区域副标题" path="title" :show-require-mark="false">
                  <n-input v-model:value="formValue.recommendLinkSubTitle" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('foodIndex') !== -1">
                <n-form-item label="餐厅首页推荐标题" path="title">
                  <n-input v-model:value="formValue.foodIndexTitle" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('spaIndex') !== -1">
                <n-form-item label="按摩首页推荐标题" path="title">
                  <n-input v-model:value="formValue.spaIndexTitle" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('carIndex') !== -1">
                <n-form-item label="接送机首页推荐标题" path="title">
                  <n-input v-model:value="formValue.carIndexTitle" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('outLink') !== -1">
                <n-form-item label="外链" path="outLinks" :show-require-mark="true">
                  <n-button type="primary" @click="addOutLink()">添加外链</n-button>
                </n-form-item>
                <n-table style="margin-bottom: 25px">
                  <thead>
                  <tr>
                    <th>外链标题</th>
                    <th>外链链接</th>
                    <th>按钮文字</th>
                    <th>跳转方式</th>
                    <th>操作</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr v-for="(item, index) of outLinkArr" :key="index">
                    <td>
                      <n-input v-model:value="item.title" placeholder="请输入外链标题" style="width: 150px"/>
                    </td>
                    <td>
                      <n-input v-model:value="item.link" placeholder="请输入外链链接" style="width: 250px"/>
                    </td>
                    <td>
                      <n-input v-model:value="item.buttonTxt" placeholder="请输入按钮文字" style="width: 120px"/>
                    </td>
                    <td>
                      <n-select v-model:value="item.openType" :options="outLinkOpenTypeOptions" style="width: 120px"/>
                    </td>
                    <td>
                      <n-button type="error" @click="handleDeleteOutLink(index)">删除</n-button>
                    </td>
                  </tr>
                  </tbody>
                </n-table>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('coupon') !== -1">
                <n-form-item label="优惠券" path="mchIds" :show-require-mark="true">
                  <n-button type="primary" @click="chooseCoupon()">添加优惠券</n-button>
                </n-form-item>
                <n-table style="margin-bottom: 25px">
                  <thead>
                    <tr>
                      <th>优惠券名称</th>
                      <!--                    <th>优惠券类型</th>
                    <th>优惠券金额/折扣</th>
                    <th>满多少元使用</th>-->
                      <th>适用场景</th>
                      <th>有效期</th>
                      <th>每人领取数量(>0)</th>
                      <th>操作</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(item, index) of couponArr" :key="index">
                      <td>{{ item.couponName }}</td>
                      <!--                    <td>
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
                    <td>{{ item.atLeast }}</td>-->
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
                        <n-input-number v-model:value="item.num" min="1" placeholder="每人领取数量" :show-button="false"
                          style="width: 80px" />
                      </td>
                      <td>
                        <n-button type="error" @click="handleDeleteCoupon(item.id)">删除</n-button>
                      </td>
                    </tr>
                  </tbody>
                </n-table>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('thCoupon') !== -1">
                <n-form-item label="礼品券" path="mchIds" :show-require-mark="true">
                  <n-button type="primary" @click="chooseThCoupon()">添加礼品券</n-button>
                </n-form-item>
                <n-table style="margin-bottom: 25px">
                  <thead>
                    <tr>
                      <th>礼品券名称</th>
                      <th>适用商户</th>
                      <th>领取规则</th>
                      <th>每人领取数量(>0)</th>
                      <th>有效期(天) <div style="color:grey; font-size: 12px;line-height: 12px">0表示领取当日有效</div></th>
                      <th>操作</th>
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
                        <n-input-group>
                          <n-input-number v-model:value="item.limitDays" min="1" placeholder="天数" :show-button="false" style="width: 50px" />
                          <n-input-group-label>天内可领</n-input-group-label>
                          <n-input-number v-model:value="item.perNum" min="1" placeholder="数量" :show-button="false" style="width: 50px" />
                          <n-input-group-label>张</n-input-group-label>
                        </n-input-group>
                      </td>
                      <td>
                        <n-input-number v-model:value="item.num" min="1" placeholder="每人领取数量" :show-button="false"
                          style="width: 80px" />
                      </td>
                      <td>
                        <n-input-group>
                          <n-input-number v-model:value="item.fixedTerm" min="0" placeholder="天数" :show-button="false" style="width: 50px" />
                          <n-input-group-label>天内有效</n-input-group-label>
                        </n-input-group>
<!--                        <div style="color:grey; font-size: 12px;line-height: 34px;margin: 0 0 24px">0表示领取当日有效</div>-->
                      </td>
                      <td>
                        <n-button type="error" @click="handleDeleteThCoupon(item.id)">删除</n-button>
                      </td>
                    </tr>
                  </tbody>
                </n-table>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('thCoupon') !== -1">
                <!-- 领取限制，限制周六、周日不可领取、分开勾选 -->
                <n-form-item label="礼品券领取限制" path="saturdayLimit" :show-require-mark="false">
                  <n-space>
                    <n-checkbox v-model:checked="saturdayLimitChecked">
                      周六不允许领取
                    </n-checkbox>
                    <n-checkbox v-model:checked="sundayLimitChecked">
                      周日不允许领取
                    </n-checkbox>
                  </n-space>
                </n-form-item>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('hotelDetail') !== -1">
                <n-form-item label="民宿" path="hotelIds" :show-require-mark="true">
                  <n-button type="primary" @click="chooseHotel()">添加民宿</n-button>
                </n-form-item>
                <n-table style="margin-bottom: 25px">
                  <thead>
                    <tr>
                      <th>民宿名称</th>
                      <th>操作</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(item, index) of hotelArr" :key="index">
                      <td>{{ item.name }}</td>
                      <td>
                        <n-button type="error" @click="handleDeleteHotel(item.id)">删除</n-button>
                      </td>
                    </tr>
                  </tbody>
                </n-table>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('foodDetail') !== -1">
                <n-form-item label="餐厅" path="mchIds" :show-require-mark="true">
                  <n-button type="primary" @click="chooseRestaurant()">添加餐厅</n-button>
                </n-form-item>
                <n-table style="margin-bottom: 25px">
                  <thead>
                    <tr>
                      <th>餐厅名称</th>
                      <th>营业类型</th>
                      <th>地址</th>
                      <th>操作</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(item, index) of restaurantArr" :key="index">
                      <td>{{ item.name }}</td>
                      <td>{{ item.cooperateTypeId > 0 ? item.cooperateTypeDetail.typeName : '--' }}</td>
                      <td>{{ item.detailAddress }}</td>
                      <td>
                        <n-button type="error" @click="handleDeleteRestaurant(item.id)">删除</n-button>
                      </td>
                    </tr>
                  </tbody>
                </n-table>
              </n-gi>
              <n-gi v-if="recommendLinkTypeArr.indexOf('spaDetail') !== -1">
                <n-form-item label="按摩服务" path="mchIds" :show-require-mark="true">
                  <n-button type="primary" @click="chooseSpaService()">添加按摩服务</n-button>
                </n-form-item>
                <n-table style="margin-bottom: 25px">
                  <thead>
                    <tr>
                      <th>服务名称</th>
                      <th>所属服务商</th>
                      <th>操作</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(item, index) of spaServiceArr" :key="index">
                      <td>{{ item.name }}</td>
                      <td>{{ item.ispId > 0 ? item.spaIspDetail.name : '--' }}</td>
                      <td>
                        <n-button type="error" @click="handleDeleteSpaService(item.id)">删除</n-button>
                      </td>
                    </tr>
                  </tbody>
                </n-table>
              </n-gi>
              </template>
            </n-grid>
          </n-form>
        </n-spin>

      </n-drawer-content>
    </n-drawer>
    <!--选择礼品券-->
    <ThChooseThCoupon ref="chooseThCouponRef" @reloadThCoupon="chooseThCouponInfo" />
    <!--选择优惠券-->
    <ThChooseCoupon ref="chooseCouponRef" @reloadCoupon="chooseCouponInfo" />
    <!--选餐厅-->
    <ThChooseRestaurant ref="chooseRestaurantRef" @reloadRestaurant="chooseRestaurantInfo" />
    <!--选按摩服务-->
    <ThChooseSpaService ref="chooseSpaServiceRef" @reloadSpaService="chooseSpaServiceInfo" />
    <!--选物业-->
    <ThChooseHotel ref="chooseHotelRef" @reloadHotel="chooseHotelInfo" />
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, nextTick } from 'vue';
import { Edit, View } from '@/api/homepageArticles';
import { newState, options, rules, State } from './model';
import Editor from '@/components/Editor/editor.vue';
import { NButton, useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';
import { subDays } from "date-fns/esm";
import ThChooseThCoupon from "./chooseThCoupon.vue";
import ThChooseCoupon from "./chooseCoupon.vue";
import ThChooseRestaurant from "./chooseRestaurant.vue";
import ThChooseSpaService from "./chooseSpaService.vue";
import ThChooseHotel from "./chooseHotel.vue";
import { ThCouponMchInter } from "/#/thCoupon";
import { All as ThCouponAll } from "@/api/thCoupon";
import { All } from "@/api/pmsCouponType";
import { All as RestaurantAll } from "@/api/foodRestaurant";
import { All as SpaServiceAll } from "@/api/spaService";
import { All as HotelAll } from "@/api/pmsProperty";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});
const recommendLinkTypeArr = ref([]);

// ---礼品券---
interface SelectThCouponInter {
  id: number;
  couponName: string;
  mchList: ThCouponMchInter[];
  num: number;
  perNum: number;
  limitDays: number;
  fixedTerm: number;
}

interface ThCouponIdsNumInter {
  couponId: number;
  num: number;
  perNum: number;
  limitDays: number;
  fixedTerm: number;
}
const thCouponIdsNumArr = ref<ThCouponIdsNumInter[]>([]);
const thCouponArr = ref<SelectThCouponInter[]>([])
const chooseThCouponRef = ref();
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
const chooseCouponRef = ref();
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
const chooseRestaurantRef = ref();
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
const chooseSpaServiceRef = ref();
// ---按摩-END---

// ---酒店---
interface SelectHotelInter {
  id: number;
  name: string;
  close: number;
}
const hotelIdsArr = ref([]);
const hotelArr = ref<SelectHotelInter[]>([])
const chooseHotelRef = ref();
// ---酒店-END---

// ---外链---
interface OutLinkInter {
  title: string;
  link: string;
  buttonTxt: string;
  openType: number;
}
const outLinkArr = ref<OutLinkInter[]>([])
const outLinkOpenTypeOptions = ref([
  { label: '内部webview', value: 1 },
  { label: '外部浏览器', value: 2 }
]);

// 添加计算属性确保复选框状态正确
const saturdayLimitChecked = computed({
  get: () => {
    const limits = formValue.value.limitWeek ? formValue.value.limitWeek.split(',') : [];
    return limits.includes('6');
  },
  set: (value) => {
    const limits = formValue.value.limitWeek ? formValue.value.limitWeek.split(',') : [];
    if (value) {
      // 不允许周六 - 添加6
      if (!limits.includes('6')) {
        limits.push('6');
      }
      formValue.value.limitWeek = limits.join(',');
    } else {
      // 允许周六 - 移除6
      formValue.value.limitWeek = limits.filter(limit => limit !== '6').join(',');
    }
  }
});

const sundayLimitChecked = computed({
  get: () => {
    const limits = formValue.value.limitWeek ? formValue.value.limitWeek.split(',') : [];
    return limits.includes('7');
  },
  set: (value) => {
    const limits = formValue.value.limitWeek ? formValue.value.limitWeek.split(',') : [];
    if (value) {
      // 不允许周日 - 添加7
      if (!limits.includes('7')) {
        limits.push('7');
      }
      formValue.value.limitWeek = limits.join(',');
    } else {
      // 允许周日 - 移除7
      formValue.value.limitWeek = limits.filter(limit => limit !== '7').join(',');
    }
  }
});
// ---外链-END---

async function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    recommendLinkTypeArr.value = []
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
    outLinkArr.value = []

    return;
  }

  // 编辑
  loading.value = true;

  await getInfo(state.id)

  loading.value = false;
}

async function getInfo(id) {
  const res = await View({ id: id, isLanguage: true });
  formValue.value = res;

  // 强制设置默认值（临时解决方案，数据库添加字段后可移除）
  // 使用Object.assign确保响应式
  // Object.assign(formValue.value, {
  //   limitWeek: ''
  // });

  // console.log('limitWeek:', formValue.value.limitWeek);

  recommendLinkTypeArr.value = res.recommendLinkType.split(',')

  console.log(' recommendLinkTypeArr.value', recommendLinkTypeArr.value)

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

  outLinkArr.value = []

  if (res.couponList) {
    res.couponList.forEach((item) => {
      if (item.couponType == 'thCoupon') {
        thCouponIdsNumArr.value.push({ couponId: item.couponId, num: item.availableQuantity, perNum: item.perDayAvailable, limitDays: item.limitDays || 1, fixedTerm: item.fixedTerm || 0 })
      } else {
        couponIdsNumArr.value.push({ couponId: item.couponId, num: item.availableQuantity })
      }
    })
  }
  if (thCouponIdsNumArr.value.length > 0) {
    await getThCouponList()
    if (thCouponArr.value.length > 0) {
      thCouponArr.value.forEach(m => {
        m.num = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.num || 1;
        m.perNum = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.perNum || 1;
        m.limitDays = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.limitDays || 1;
        m.fixedTerm = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.fixedTerm || 0
      })
    }
  }
  if (couponIdsNumArr.value.length > 0) {
    await getCouponList()
    if (couponArr.value.length > 0) {
      couponArr.value.forEach(m => {
        m.num = couponIdsNumArr.value.find(n => n.couponId == m.id)?.num || 1
      })
    }
  }

  if (res.restaurantIds) {
    restaurantIdsArr.value = res.restaurantIds.split(',').map((item) => {
      return Number(item)
    })
  } else {
    restaurantIdsArr.value = []
  }
  if (restaurantIdsArr.value.length > 0) {
    await getRestaurantList()
  }
  // 按摩服务
  if (res.spaServiceIds) {
    spaServiceIdsArr.value = res.spaServiceIds.split(',').map((item) => {
      return Number(item)
    })
  } else {
    spaServiceIdsArr.value = []
  }
  if (spaServiceIdsArr.value.length > 0) {
    await getSpaServiceList()
  }
  // 酒店
  if (res.hotelIds) {
    hotelIdsArr.value = res.hotelIds.split(',').map((item) => {
      return Number(item)
    })
  } else {
    hotelIdsArr.value = []
  }
  if (hotelIdsArr.value.length > 0) {
    await getHotelList()
  }

  // 外链
  if (res.outLinks && res.outLinks.length > 0) {
    outLinkArr.value = res.outLinks.map(item => ({
      title: item.title || '',
      link: item.link || '',
      buttonTxt: item.buttonTxt || '',
      openType: item.openType || 1
    }))
  } else {
    outLinkArr.value = []
  }
}

function confirmForm(e) {
  if (!formValue.value.no){
    formBtnLoading.value = false;
    message.error('请填写活动编号');
    return false;
  }

  if (formValue.value.recommendLinkType.indexOf('coupon') !== -1) {
    if (couponArr.value.length <= 0) {
      formBtnLoading.value = false;
      message.error('请选择优惠券');
      return false;
    }
    formValue.value.couponsArr = couponArr.value.map(m => ({
      couponId: m.id,
      availableQuantity: m.num
    }))
  }
  if (formValue.value.recommendLinkType.indexOf('thCoupon') !== -1) {
    if (thCouponArr.value.length <= 0) {
      formBtnLoading.value = false;
      message.error('请选择礼品券');
      return false;
    }
    for (const m of thCouponArr.value) {
      if (!Number.isInteger(m.limitDays) || m.limitDays < 1) {
        formBtnLoading.value = false;
        message.error(`礼品券【${m.couponName}】的限制天数必须为大于0的整数`);
        return false;
      }
      if (!Number.isInteger(m.perNum) || m.perNum < 1) {
        formBtnLoading.value = false;
        message.error(`礼品券【${m.couponName}】的天数内可领取数量必须为大于0的整数`);
        return false;
      }
    }
    formValue.value.thCouponsArr = thCouponArr.value.map(m => ({
      couponId: m.id,
      availableQuantity: m.num,
      perDayAvailable: m.perNum,
      limitDays: m.limitDays,
      fixedTerm: m.fixedTerm || 0,
    }))
  }

  if (formValue.value.recommendLinkType.indexOf('foodDetail') !== -1) {
    if (restaurantIdsArr.value.length <= 0) {
      formBtnLoading.value = false;
      message.error('请选择餐厅');
      return false;
    }
    formValue.value.restaurantIds = restaurantIdsArr.value.join(',')
  } else {
    formValue.value.restaurantIds = "";
  }

  if (formValue.value.recommendLinkType.indexOf('spaDetail') !== -1) {
    if (spaServiceIdsArr.value.length <= 0) {
      formBtnLoading.value = false;
      message.error('请选择按摩服务');
      return false;
    }
    formValue.value.spaServiceIds = spaServiceIdsArr.value.join(',')
  } else {
    formValue.value.spaServiceIds = "";
  }

  if (formValue.value.recommendLinkType.indexOf('hotelDetail') !== -1) {
    if (hotelIdsArr.value.length <= 0) {
      formBtnLoading.value = false;
      message.error('请选择酒店');
      return false;
    }
    formValue.value.hotelIds = hotelIdsArr.value.join(',')
  } else {
    formValue.value.hotelIds = "";
  }

  if (formValue.value.recommendLinkType.indexOf('foodIndex') !== -1) {
    if (!formValue.value.foodIndexTitle || formValue.value.foodIndexTitle == "") {
      formBtnLoading.value = false;
      message.error('请输入餐厅推荐标题');
      return false;
    }
  }
  if (formValue.value.recommendLinkType.indexOf('spaIndex') !== -1) {
    if (!formValue.value.spaIndexTitle || formValue.value.spaIndexTitle == "") {
      formBtnLoading.value = false;
      message.error('请输入按摩推荐标题');
      return false;
    }
  }
  if (formValue.value.recommendLinkType.indexOf('carIndex') !== -1) {
    if (!formValue.value.carIndexTitle || formValue.value.carIndexTitle == "") {
      formBtnLoading.value = false;
      message.error('请输入接送机推荐标题');
      return false;
    }
  }

  if (formValue.value.recommendLinkType.indexOf('outLink') !== -1) {
    if (outLinkArr.value.length <= 0) {
      formBtnLoading.value = false;
      message.error('请至少添加一个外链');
      return false;
    }
    // 验证每个外链的必填项
    for (let i = 0; i < outLinkArr.value.length; i++) {
      const link = outLinkArr.value[i];
      if (!link.title || link.title.trim() === '') {
        formBtnLoading.value = false;
        message.error(`第${i + 1}个外链的标题不能为空`);
        return false;
      }
      if (!link.link || link.link.trim() === '') {
        formBtnLoading.value = false;
        message.error(`第${i + 1}个外链的链接不能为空`);
        return false;
      }
      // 验证外链URL格式
      const urlPattern = /^https?:\/\/.+/;
      if (!urlPattern.test(link.link)) {
        formBtnLoading.value = false;
        message.error(`第${i + 1}个外链地址格式不正确，必须以http://或https://开头`);
        return false;
      }
      if (!link.buttonTxt || link.buttonTxt.trim() === '') {
        formBtnLoading.value = false;
        message.error(`第${i + 1}个外链的按钮文字不能为空`);
        return false;
      }
    }
    // 将外链数组赋值给 formValue
    formValue.value.outLinks = outLinkArr.value
  } else {
    formValue.value.outLinks = []
  }

  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          formBtnLoading.value = false;
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        console.log(err);
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

function disablePreviousDate(ts: number) {
  return ts < subDays(Date.now(), 1).valueOf();
}

function handleUpdateValue(value) {
  value = value.filter(item => item !== null && item !== undefined && item !== '')
  formValue.value.recommendLinkType = value.join(',')
}

// ==============================================礼品券=====================================================
// 选择礼品券相关
function chooseThCoupon() {
  chooseThCouponRef.value.openModal(thCouponIdsNumArr.value.map(m => m.couponId));
}
// 删除礼品券
function handleDeleteThCoupon(id) {
  thCouponArr.value = thCouponArr.value.filter(m => m.id !== id)
  thCouponIdsNumArr.value = thCouponIdsNumArr.value.filter(m => m.couponId !== id)
}

// 礼品券选择回调
async function chooseThCouponInfo(selectedIds) {
  thCouponIdsNumArr.value = selectedIds.map(m => ({ couponId: m, num: 1, perNum: 1, limitDays: 1, fixedTerm: 0 }))
  loading.value = true;
  const oldMap = new Map(thCouponArr.value.map(m => [m.id, m]))
  console.log('oldMap', oldMap)
  await getThCouponList()

  thCouponArr.value = thCouponArr.value.map(m => ({
    ...m,
    num: oldMap.get(m.id)?.num || 1,
    perNum: oldMap.get(m.id)?.perNum || 1,
    limitDays: oldMap.get(m.id)?.limitDays || 1,
    fixedTerm: oldMap.get(m.id)?.fixedTerm || 0
  }))
  console.log('thCouponArr', thCouponArr.value)
  loading.value = false;
}

// 获取礼品券列表
async function getThCouponList() {
  const res = await ThCouponAll({
    couponIds: thCouponIdsNumArr.value.map(m => m.couponId).join(',')
  });
  thCouponArr.value = res.list
  console.log('thCouponArr', thCouponArr.value)
}

// ==============================================优惠券=====================================================
// 选择优惠券相关
function chooseCoupon() {
  chooseCouponRef.value.openModal(couponIdsNumArr.value.map(m => m.couponId));
}
// 删除礼品券
function handleDeleteCoupon(id) {
  couponArr.value = couponArr.value.filter(m => m.id !== id)
  couponIdsNumArr.value = couponIdsNumArr.value.filter(m => m.couponId !== id)
}

// 礼品券选择回调
async function chooseCouponInfo(selectedIds) {
  couponIdsNumArr.value = selectedIds.map(m => ({ couponId: m }))
  loading.value = true;
  const oldMap = new Map(couponArr.value.map(m => [m.id, m]))
  console.log('couponArr_oldMap', oldMap)
  await getCouponList()
  couponArr.value = couponArr.value.map(m => ({
    ...m,
    num: oldMap.get(m.id)?.num || 1
  }))
  console.log('couponArr', couponArr.value)
  loading.value = false;
}

// 获取优惠券列表
async function getCouponList() {
  const res = await All({
    couponIds: couponIdsNumArr.value.map(m => m.couponId).join(',')
  });
  couponArr.value = res.list
  console.log('couponArr', couponArr.value)
}
// ==============================================优惠券END=====================================================

// ==============================================餐厅=====================================================
// 选择餐厅相关
function chooseRestaurant() {
  chooseRestaurantRef.value.openModal(restaurantIdsArr.value);
}
// 删除餐厅
function handleDeleteRestaurant(id) {
  restaurantArr.value = restaurantArr.value.filter(m => m.id !== id)
  restaurantIdsArr.value = restaurantIdsArr.value.filter(m => m !== id)
}

// 餐厅选择回调
async function chooseRestaurantInfo(selectedIds) {
  restaurantIdsArr.value = selectedIds
  loading.value = true;
  await getRestaurantList()
  loading.value = false;
}

// 获取餐厅列表
async function getRestaurantList() {
  const res = await RestaurantAll({
    restaurantIds: restaurantIdsArr.value.join(',')
  });
  restaurantArr.value = res.list
  console.log('restaurantArr', restaurantArr.value)
}
// ==============================================餐厅END=====================================================

// ==============================================按摩服务=====================================================
// 选择按摩相关
function chooseSpaService() {
  chooseSpaServiceRef.value.openModal(spaServiceIdsArr.value);
}
// 删除按摩
function handleDeleteSpaService(id) {
  spaServiceArr.value = spaServiceArr.value.filter(m => m.id !== id)
  spaServiceIdsArr.value = spaServiceIdsArr.value.filter(m => m !== id)
}

// 按摩选择回调
async function chooseSpaServiceInfo(selectedIds) {
  spaServiceIdsArr.value = selectedIds
  loading.value = true;
  await getSpaServiceList()
  loading.value = false;
}

// 获取按摩列表
async function getSpaServiceList() {
  const res = await SpaServiceAll({
    serviceIds: spaServiceIdsArr.value.join(',')
  });
  spaServiceArr.value = res.list
  console.log('spaServiceArr', spaServiceArr.value)
}
// ==============================================按摩END=====================================================

// ==============================================酒店=====================================================
// 选择酒店相关
function chooseHotel() {
  chooseHotelRef.value.openModal(hotelIdsArr.value);
}
// 删除酒店
function handleDeleteHotel(id) {
  hotelArr.value = hotelArr.value.filter(m => m.id !== id)
  hotelIdsArr.value = hotelIdsArr.value.filter(m => m !== id)
}

// 酒店选择回调
async function chooseHotelInfo(selectedIds) {
  hotelIdsArr.value = selectedIds
  loading.value = true;
  await getHotelList()
  loading.value = false;
}

// 获取酒店列表
async function getHotelList() {
  const res = await HotelAll({
    ids: hotelIdsArr.value.join(',')
  });
  hotelArr.value = res.list
  console.log('hotelArr', hotelArr.value)
}
// ==============================================酒店END=====================================================

// ==============================================外链=====================================================
// 添加外链
function addOutLink() {
  outLinkArr.value.push({
    title: '',
    link: '',
    buttonTxt: '',
    openType: 1
  })
}

// 删除外链
function handleDeleteOutLink(index) {
  outLinkArr.value.splice(index, 1)
}
// ==============================================外链END=====================================================

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>
