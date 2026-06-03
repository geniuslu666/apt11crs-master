<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
        :label-placement="settingStore.isMobile ? 'top' : 'left'"
        :label-width="150"
        class="py-4"
      >
        <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
          <template v-if="formValue.cooperateTypeId != 4">
          <n-gi span="1">
            <n-form-item label="预约日期" path="orderTimeType">
              <n-radio-group v-model:value="formValue.orderTimeType" name="orderTimeType">
                <n-space>
                  <n-radio :value="1">
                    每天
                  </n-radio>
                  <n-radio :value="2">
                    自定义
                  </n-radio>
                </n-space>
              </n-radio-group>
            </n-form-item>
            <n-form-item label=" " path="orderTimeWeekArr" v-if="formValue.orderTimeType == 2">
              <n-checkbox-group v-model:value="formValue.orderTimeWeekArr">
                <n-space>
                  <n-checkbox
                    v-for="item in weekList"
                    :key="item.value"
                    :value="item.value"
                    :label="item.label"
                  />
                </n-space>
              </n-checkbox-group>
            </n-form-item>
          </n-gi>
<!--          <n-gi span="1">-->
<!--            <n-form-item label="定休日" path="restTimeWeekArr">-->
<!--              <n-checkbox-group v-model:value="formValue.restTimeWeekArr">-->
<!--                <n-space>-->
<!--                  <n-checkbox-->
<!--                    v-for="item in weekList"-->
<!--                    :key="item.value"-->
<!--                    :value="item.value"-->
<!--                    :label="item.label"-->
<!--                  />-->
<!--                </n-space>-->
<!--              </n-checkbox-group>-->
<!--            </n-form-item>-->
<!--            <n-form-item label=" " :path="'restTimeDate'+index" v-for="(item, index) in restDateFormArr" :key="index">-->
<!--              <n-date-picker placeholder="请选择自定义日期" v-model:formatted-value="item.date" value-format="yyyy-MM-dd" type="date" />-->
<!--              <n-button text type="success" style="margin-left: 12px" attr-type="button" @click="addItem" v-if="index == 0">-->
<!--                <template #icon>-->
<!--                  <n-icon>-->
<!--                    <PlusCircleOutlined />-->
<!--                  </n-icon>-->
<!--                </template>-->
<!--              </n-button>-->
<!--              <n-button text type="error" style="margin-left: 12px" attr-type="button" @click="removeItem(index)" v-if="index > 0">-->
<!--                <template #icon>-->
<!--                  <n-icon>-->
<!--                    <MinusCircleOutlined />-->
<!--                  </n-icon>-->
<!--                </template>-->
<!--              </n-button>-->
<!--            </n-form-item>-->
<!--          </n-gi>-->
          <n-gi span="1" v-if="formValue.orderConfirmDayType == 1">
            <n-form-item label="预约确定模式" path="orderConfirm">
              <n-input-group>
<!--                <n-input-group-label>预定</n-input-group-label>-->
<!--                <n-input-number placeholder="请输入" :min="0" :precision="0" v-model:value="formValue.orderConfirmBeforeDays" style="width: 100px" />-->
<!--                <n-input-group-label>天内手动确认，预计确认时长</n-input-group-label>-->
                <n-input-group-label>手动确认，预计确认时长</n-input-group-label>
                <n-input-number placeholder="请输入确认时长" :min="0" :precision="0" v-model:value="formValue.orderConfirmTime" style="width: 100px" />
                <n-input-group-label>分钟</n-input-group-label>
              </n-input-group>
            </n-form-item>
<!--            <n-form-item label=" " path="orderConfirm">-->
<!--              <n-input-group>-->
<!--                <n-input-group-label>大于{{ formValue.orderConfirmBeforeDays }}天自动确认</n-input-group-label>-->
<!--              </n-input-group>-->
<!--            </n-form-item>-->
          </n-gi>
          <n-gi span="1" v-if="formValue.orderConfirmDayType == 2">
            <n-form-item label="预约确定模式" path="orderConfirm">
              <n-input-group>
                <n-input-group-label>预定</n-input-group-label>
                <n-input-number placeholder="请输入" :min="0" :precision="0" v-model:value="formValue.orderConfirmBeforeDays" style="width: 100px" />
                <n-input-group-label>天前自动确认</n-input-group-label>
              </n-input-group>
            </n-form-item>
            <n-form-item label=" " path="orderConfirm">
              <n-input-group>
                <n-input-group-label>{{ formValue.orderConfirmBeforeDays }}天之后手动确认，预计确认时长</n-input-group-label>
                <n-input-number placeholder="请输入确认时长" :min="0" :precision="0" v-model:value="formValue.orderConfirmTime" style="width: 100px" />
                <n-input-group-label>分钟</n-input-group-label>
              </n-input-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="预约日期" path="openTimeType">
              <n-radio-group v-model:value="formValue.openTimeType" name="openTimeType">
                <n-space>
                  <n-radio :value="1">
                    按时段
                  </n-radio>
                  <n-radio :value="2">
                    按时间点
                  </n-radio>
                </n-space>
              </n-radio-group>
            </n-form-item>
            <n-form-item label="预约时段" path="amTime" v-if="formValue.openTimeType == 1">
              早市：
              <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.amOpenTime" />
              <span style="margin: 0 5px">~</span>
              <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.amCloseTime" />
            </n-form-item>
            <n-form-item label=" " path="pmTime" v-if="formValue.openTimeType == 1">
              晚市：
              <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.pmOpenTime" />
              <span style="margin: 0 5px">~</span>
              <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.pmCloseTime" />
            </n-form-item>
            <n-form-item label="时段间隔" path="timeDuration" v-if="formValue.openTimeType == 1">
              <n-radio-group v-model:value="formValue.timeDuration" name="timeDuration">
                <n-radio-button
                  v-for="item in timeDurationArr"
                  :key="item.value"
                  :value="item.value"
                  :label="item.label"
                />
              </n-radio-group>
              <n-input-group-label>分钟</n-input-group-label>
            </n-form-item>
            <n-form-item label="时间点" :path="'timePoints'+index" v-for="(item, index) in timePointsFormArr" :key="index" v-if="formValue.openTimeType == 2">
              <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="item.time" />
              <n-button text type="success" style="margin-left: 12px" attr-type="button" @click="addTimeItem" v-if="index == 0">
                <template #icon>
                  <n-icon>
                    <PlusCircleOutlined />
                  </n-icon>
                </template>
              </n-button>
              <n-button text type="error" style="margin-left: 12px" attr-type="button" @click="removeTimeItem(index)" v-if="index > 0">
                <template #icon>
                  <n-icon>
                    <MinusCircleOutlined />
                  </n-icon>
                </template>
              </n-button>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="最大容纳预定数" path="orderMaxType">
              <n-radio-group v-model:value="formValue.orderMaxType" name="orderMaxType">
                <n-space>
                  <n-radio :value="1">
                    按时段
                  </n-radio>
                  <n-radio :value="2">
                    全局
                  </n-radio>
                </n-space>
              </n-radio-group>
            </n-form-item>
            <n-form-item :label="formValue.orderMaxType == 1 ? '时段最大容纳预定数' : '全局最大容纳预定数'" path="timeDurationMax" style="margin-bottom: 24px">
              <n-input-group>
                <n-input-number placeholder="请输入预定数" :min="0" :precision="0" v-model:value="formValue.timeDurationMax" style="width: 100px" />
                <n-input-group-label>人</n-input-group-label>
              </n-input-group>
              <template #feedback>
                填0则不限制
              </template>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="提前预约" path="advanceOrderDay" style="margin-bottom: 24px">
              <n-input-group>
                <n-input-group-label>需至少提前</n-input-group-label>
                <n-input-number placeholder="请输入天数" :min="1" :precision="0" v-model:value="formValue.advanceOrderDay" style="width: 100px" />
                <n-input-group-label>天</n-input-group-label>
              </n-input-group>
<!--              <template #feedback>-->
<!--                填0则当日可约-->
<!--              </template>-->
            </n-form-item>
          </n-gi>
          <n-gi span="1" v-if="formValue.orderTimeType == 1">
            <n-form-item label="预约处理截止时间" path="maxOrderDay">
              <n-input-group>
                <n-input-group-label>当天预约处理截止时间</n-input-group-label>
                <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.dayTimeLimit" />
<!--                <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.dayTimeLimit" />-->
                <n-input-group-label>前</n-input-group-label>
              </n-input-group>
            </n-form-item>
          </n-gi>
          <n-gi span="1">
            <n-form-item label="最长预约" path="maxOrderDay">
              <n-input-group>
                <n-input-group-label>最长可预约</n-input-group-label>
                <n-input-number placeholder="请输入天数" :min="0" :precision="0" :show-button="false" v-model:value="formValue.maxOrderDay" style="width: 100px" />
                <n-input-group-label>天后</n-input-group-label>
              </n-input-group>
            </n-form-item>
          </n-gi>

            <n-gi span="1">
              <n-form-item label="预定模式" path="orderTimeType">
                <n-radio-group v-model:value="formValue.orderMode" name="orderTimeType">
                  <n-space>
                    <n-radio :value="'ALLAMOUNT'">
                      全款模式
                    </n-radio>
                    <n-radio :value="'DEPOSIT'">
                      定金模式
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
            </n-gi>

            <n-gi span="1" v-if="formValue.orderMode == 'DEPOSIT'">
              <n-form-item label="定金支付比例" path="depositRate">
                <n-input-group>
                  <n-input-number placeholder="请输入" :min="1" :max="99" v-model:value="formValue.depositRate" :show-button="false" style="width: 80px"/>
                  <n-input-group-label>%</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
          </template>
          <template v-else>
            <n-gi span="1">
              <n-form-item label="预约日期" path="orderTimeType">
                <n-radio-group v-model:value="formValue.orderTimeType" name="orderTimeType">
                  <n-space>
                    <n-radio :value="1">
                      每天
                    </n-radio>
                    <n-radio :value="2">
                      自定义
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
              <n-form-item label=" " path="orderTimeWeekArr" v-if="formValue.orderTimeType == 2">
                <n-checkbox-group v-model:value="formValue.orderTimeWeekArr">
                  <n-space>
                    <n-checkbox
                      v-for="item in weekList"
                      :key="item.value"
                      :value="item.value"
                      :label="item.label"
                    />
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item :label="formValue.orderMaxType == 1 ? '时段最大容纳预定数' : '全局最大容纳预定数'" path="timeDurationMax" style="margin-bottom: 24px">
                <n-input-group>
                  <n-input-number placeholder="请输入预定数" :min="0" :precision="0" v-model:value="formValue.timeDurationMax" style="width: 100px" />
                  <n-input-group-label>人</n-input-group-label>
                </n-input-group>
                <template #feedback>
                  填0则不限制
                </template>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="提前预约" path="advanceOrderDay" style="margin-bottom: 24px">
                <n-input-group>
                  <n-input-group-label>需至少提前</n-input-group-label>
                  <n-input-number placeholder="请输入天数" :min="1" :precision="0" v-model:value="formValue.advanceOrderDay" style="width: 100px" />
                  <n-input-group-label>天</n-input-group-label>
                </n-input-group>
                <!--              <template #feedback>-->
                <!--                填0则当日可约-->
                <!--              </template>-->
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预约处理截止时间" path="maxOrderDay">
                <n-input-group>
                  <n-input-group-label>当天预约处理截止时间</n-input-group-label>
                    <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.dayTimeLimit" />
                  <n-input-group-label>前</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="最长预约" path="maxOrderDay">
                <n-input-group>
                  <n-input-group-label>最长可预约</n-input-group-label>
                  <n-input-number placeholder="请输入天数" :min="0" :precision="0" :show-button="false" v-model:value="formValue.maxOrderDay" style="width: 100px" />
                  <n-input-group-label>天后</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>

            <n-gi span="1">
              <n-form-item label="是否允许取消" path="toretaCancelEnable">
                <n-input-group style="margin-left: 8px">
                  <n-switch v-model:value="formValue.toretaCancelEnable" :unchecked-value="2" :checked-value="1">
                    <template #checked>
                      允许
                    </template>
                    <template #unchecked>
                      不允许
                    </template>
                  </n-switch>
                </n-input-group>
              </n-form-item>
<!--              <n-form-item label="是否允许取消" path="toretaCancelEnable">
                <n-radio-group v-model:value="formValue.toretaCancelEnable" name="toretaCancelEnable">
                  <n-space>
                    <n-radio :value="1">
                      允许
                    </n-radio>
                    <n-radio :value="2">
                      不允许
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>-->
              <n-form-item label="允许取消时间" path="cancelPolicyOpen" v-if="formValue.toretaCancelEnable == 1">
                <n-input-group>
                  <n-input-group-label>距离到店时间</n-input-group-label>
                  <n-input-number placeholder="请输入" :min="0" v-model:value="formValue.toretaCancelLimitDay" style="width: 100px" />
                  <n-input-group-label>天</n-input-group-label>
                  <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.toretaCancelLimitTime" />
                  <n-input-group-label>点前，允许取消</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
          </template>

          <n-gi span="1">
              <n-form-item label="日期限制" path="maxDatetimeOrderOpen">
                <n-radio-group v-model:value="formValue.maxDatetimeOrderOpen" name="maxDatetimeOrderOpen">
                  <n-space>
                    <n-radio :value="1">
                      开启
                    </n-radio>
                    <n-radio :value="2">
                      关闭
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
              <n-form-item :label="index == 0 ? '日期' : ' '" :path="'maxDatetimeOrderDate'+index" v-if="formValue.maxDatetimeOrderOpen == 1" v-for="(item, index) in orderDateFormArr" :key="index">
                <n-date-picker placeholder="日期" format="yyyy-MM-dd" v-model:formatted-value="item.date"/>
                <n-button text type="success" style="margin-left: 12px" attr-type="button" @click="addDateItem" v-if="index == 0">
                  <template #icon>
                    <n-icon>
                      <PlusCircleOutlined />
                    </n-icon>
                  </template>
                </n-button>
                <n-button text type="error" style="margin-left: 12px" attr-type="button" @click="removeDateItem(index)" v-if="index > 0">
                  <template #icon>
                    <n-icon>
                      <MinusCircleOutlined />
                    </n-icon>
                  </template>
                </n-button>
              </n-form-item>
            </n-gi>
        </n-grid>
        <div style="text-align: center;">
          <n-space justify="center">
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
              确定
            </n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>
<script setup lang="ts">
import {onMounted, reactive, ref, watch} from "vue";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Edit} from "@/api/foodRestaurant";
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {useRouter} from "vue-router";
import {PlusCircleOutlined,MinusCircleOutlined} from "@vicons/antd";
import { isTemplateExpression } from "typescript";

const restDateFormArr = ref([
  {
    date: null
  }
]);
const timePointsFormArr = ref([
  {
    time: null
  }
]);
const timeDurationArr = reactive([
  {
    label: '15',
    value: 15
  },
  {
    label: '30',
    value: 30
  },
  {
    label: '60',
    value: 60
  },
  {
    label: '90',
    value: 90
  },
  {
    label: '120',
    value: 120
  },
])
const props = defineProps({
  formData: {
    type: Object || null,
    default: null,
  },
});
const emit = defineEmits(['reloadInfo']);
const router = useRouter();
const tabsViewStore = useTabsViewStore();
const show = ref(false);
const weekList = ref([
  {
    label: '周一',
    value: 1,
  },
  {
    label: '周二',
    value: 2,
  },
  {
    label: '周三',
    value: 3,
  },
  {
    label: '周四',
    value: 4,
  },
  {
    label: '周五',
    value: 5,
  },
  {
    label: '周六',
    value: 6,
  },
  {
    label: '周日',
    value: 7,
  },
])
const orderDateFormArr = ref([
  {
    date: null
  }
])
const formValue = ref({
  type: 'operation',
  id: 0,
  orderMode: '',
  cooperateTypeId: 0,
  orderTimeType: 1,
  orderTimeWeekArr: [],
  orderTimeWeek: '',
  restTimeWeekArr: [],
  restTimeWeek: '',
  restTimeDate: '',
  orderConfirmType: 1,
  orderConfirmDayType: 1,
  orderConfirmBeforeDays: 0,
  orderConfirmTime: 0,
  openTimeType: 1,
  amOpenTime: null,
  amCloseTime: null,
  pmOpenTime: null,
  pmCloseTime: null,
  timePoints: '',
  timePointsArr: [],
  timeDuration: 15,
  orderMaxType: 1,
  timeDurationMax: 0,
  advanceOrderDay: 0,
  dayTimeLimit: null, // 当天预约处理截止时间
  maxOrderDay: 0,
  depositRate: 0,
  toretaCancelEnable: 2, // Toreta是否允许取消 1 允许  2 不允许
  toretaCancelLimitDay: 0, // Toreta允许取消几天前
  toretaCancelLimitTime: null, // Toreta允许取消几天前
  maxDatetimeOrderOpen: 2,
  maxDatetimeOrderDate: '',
});
const settingStore = useProjectSettingStore();
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();

const rules = ref({});

const removeItem = (index: number) => {
  restDateFormArr.value.splice(index, 1)
}

const addItem = () => {
  restDateFormArr.value.push({ date: null })
}

const removeTimeItem = (index: number) => {
  timePointsFormArr.value.splice(index, 1)
}

const addTimeItem = () => {
  timePointsFormArr.value.push({ time: null })
}

const removeDateItem = (index: number) => {
  orderDateFormArr.value.splice(index, 1)
}

const addDateItem = () => {
  orderDateFormArr.value.push({ date: null })
}

function confirmForm(e) {
  let restDateList = restDateFormArr.value.map((item) => {
    return item.date
  })
  restDateList = restDateList.filter(item => item !== '' && item !== null && item !== undefined && item === item);

  let timePointsList = timePointsFormArr.value.map((item) => {
    return item.time
  })
  timePointsList = timePointsList.filter(item => item !== '' && item !== null && item !== undefined && item === item);

  let orderDateList = orderDateFormArr.value.map((item) => {
    return item.date
  })
  orderDateList = orderDateList.filter(item => item !== '' && item !== null && item !== undefined && item === item);

  if(formValue.value.cooperateTypeId != 4){
    // 如果预定模式是定金模式，那么定金比例不能为空，且大于等于1小于等于99
    if(formValue.value.orderMode == "DEPOSIT"){
      if(!formValue.value.depositRate || parseFloat(formValue.value.depositRate) <= 0){
        message.error('请设置定金支付比例');
        formBtnLoading.value = false;
        return false;
      }
      if(parseFloat(formValue.value.depositRate) > 99 || parseFloat(formValue.value.depositRate) < 1){
        message.error('请设置定金支付比例，最大99最小1');
        formBtnLoading.value = false;
        return false;
      }
    }else{
      // 全款模式则不需要设置定金比例，定金比例默认为0
      formValue.value.depositRate = 0;
    }
  }

  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      formValue.value.orderTimeWeek = formValue.value.orderTimeWeekArr.join(',')
      formValue.value.restTimeWeek = formValue.value.restTimeWeekArr.join(',')
      formValue.value.restTimeDate = restDateList.join(',')
      formValue.value.timePoints = timePointsList.join(',')
      formValue.value.maxDatetimeOrderDate = formValue.value.maxDatetimeOrderOpen == 2 ? '' : orderDateList.join(',')
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          if(formValue.value.id > 0){
            emit('reloadInfo');
          }else{
            // todo 关闭当前tab页，并跳转到等级列表页
            setTimeout(() => {
              tabsViewStore.closeSignal('2');
              router.push({ name: 'foodRestaurantIndex', params: {  } });
            }, 500);

          }
        }, 500);
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

async function initData(data){
  if(data.restTimeDate){
    let restDateList = data.restTimeDate.split(',');
    let restDateFormInitArr = [];
    restDateList.forEach((item,index) => {
      let restDateFormItem = {
        date: item
      }
      restDateFormInitArr.push(restDateFormItem)
    })
    restDateFormArr.value = restDateFormInitArr
  }
  if(data.timePoints){
    let timePointsList = data.timePoints.split(',');
    let timePointsFormInitArr = [];
    timePointsList.forEach((item,index) => {
      let timePointsFormItem = {
        time: item
      }
      timePointsFormInitArr.push(timePointsFormItem)
    })
    timePointsFormArr.value = timePointsFormInitArr
  }
  if(data.maxDatetimeOrderDate){
    let orderDateList = data.maxDatetimeOrderDate.split(',');
    let orderDateFormInitArr = [];
    orderDateList.forEach((item,index) => {
      let orderDateFormItem = {
        date: item
      }
      orderDateFormInitArr.push(orderDateFormItem)
    })
    orderDateFormArr.value = orderDateFormInitArr
  }
  formValue.value = {
    type: 'operation',
    id: data.id,
    orderMode: data.orderMode,
    cooperateTypeId: data.cooperateTypeId,
    orderTimeType: data.orderTimeType,
    orderTimeWeekArr: data.orderTimeWeek ? data.orderTimeWeek.split(',').map((item)=>{
      return parseInt(item)
    }) : [],
    orderTimeWeek: data.orderTimeWeek,
    restTimeWeekArr: data.restTimeWeek ? data.restTimeWeek.split(',').map((item)=>{
      return parseInt(item)
    }) : [],
    restTimeWeek: data.restTimeWeek,
    restTimeDate: data.restTimeDate,
    orderConfirmType: data.orderConfirmType,
    orderConfirmDayType: data.orderConfirmDayType,
    orderConfirmBeforeDays: data.orderConfirmBeforeDays,
    orderConfirmTime: data.orderConfirmTime,
    openTimeType: data.openTimeType ? data.openTimeType : 1,
    amOpenTime: data.amOpenTime ? data.amOpenTime : null,
    amCloseTime: data.amCloseTime ? data.amCloseTime : null,
    pmOpenTime: data.pmOpenTime ? data.pmOpenTime : null,
    pmCloseTime: data.pmCloseTime ? data.pmCloseTime : null,
    timePoints: data.timePoints ? data.timePoints : '',
    timePointsArr: [],
    timeDuration: data.timeDuration ? parseInt(data.timeDuration) : 15,
    orderMaxType: data.orderMaxType ? data.orderMaxType : 1,
    timeDurationMax: data.timeDurationMax,
    advanceOrderDay: data.advanceOrderDay,
    maxOrderDay: data.maxOrderDay,
    depositRate: data.depositRate,
    dayTimeLimit: data.dayTimeLimit ? data.dayTimeLimit : null, // 当天预约处理截止时间
    toretaCancelEnable: data.toretaCancelEnable ? data.toretaCancelEnable : 2, // Toreta是否允许取消 1 允许  2 不允许
    toretaCancelLimitDay: data.toretaCancelLimitDay, // Toreta允许取消几天前
    toretaCancelLimitTime: data.toretaCancelLimitTime ? data.toretaCancelLimitTime : null, // Toreta允许取消几天前
    maxDatetimeOrderOpen: data.maxDatetimeOrderOpen ? data.maxDatetimeOrderOpen : 2,
    maxDatetimeOrderDate: data.maxDatetimeOrderDate ? data.maxDatetimeOrderDate : '',
  }
}

watch(
  () => props.formData,
  async(value) => {
    await initData(value)
  }
);

onMounted(async() => {
  show.value = true;
  await initData(props.formData)

  show.value = false;
});

</script>

<style lang="less">
</style>
