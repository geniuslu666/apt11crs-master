<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">预约设置</text>
        </template>
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="150"
          class="py-4"
          style="padding-top: 0"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="预约模式" path="orderMode">
                <n-radio-group v-model:value="formValue.orderMode" name="orderMode">
                  <n-space>
                    <n-radio :value="1">
                      简易模式
                    </n-radio>
                    <n-radio :value="2">
                      普通模式
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
            </n-gi>
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
<!--            <n-gi span="1">-->
<!--              <n-form-item label="定休日" path="restTimeWeekArr">-->
<!--                <n-checkbox-group v-model:value="formValue.restTimeWeekArr">-->
<!--                  <n-space>-->
<!--                    <n-checkbox-->
<!--                      v-for="item in weekList"-->
<!--                      :key="item.value"-->
<!--                      :value="item.value"-->
<!--                      :label="item.label"-->
<!--                    />-->
<!--                  </n-space>-->
<!--                </n-checkbox-group>-->
<!--              </n-form-item>-->
<!--            </n-gi>-->
            <n-gi span="1">
              <n-form-item :label="index == 0 ? '可预约时段' : ' '" :path="'orderTime'+index" v-for="(item, index) in orderTimeFormArr" :key="index">
                <n-time-picker placeholder="开始时间" format="HH:mm" v-model:formatted-value="item.startTime" />
                <span style="margin: 0 5px">~</span>
                <n-time-picker placeholder="结束时间" format="HH:mm" v-model:formatted-value="item.endTime" />
                <n-button text type="success" style="margin-left: 12px" attr-type="button" @click="addItem" v-if="index == 0">
                  <template #icon>
                    <n-icon>
                      <PlusCircleOutlined />
                    </n-icon>
                  </template>
                </n-button>
                <n-button text type="error" style="margin-left: 12px" attr-type="button" @click="removeItem(index)" v-if="index > 0">
                  <template #icon>
                    <n-icon>
                      <MinusCircleOutlined />
                    </n-icon>
                  </template>
                </n-button>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预约时段间隔" path="timeDuration">
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
            </n-gi>
            <n-gi span="1" v-if="formValue.orderMode == 2">
              <n-form-item label="项目服务间隔" path="serviceTimeDuration" style="margin-bottom: 24px">
                <n-input-group>
                  <n-input-number placeholder="请输入" :min="0" :precision="0" :show-button="false" v-model:value="formValue.serviceTimeDuration" style="width: 100px" />
                  <n-input-group-label>分钟</n-input-group-label>
                </n-input-group>
                <template #feedback>
                  <div style="font-size: 14px; color: red">服务间隔时间是指一个项目结束后，到服务下一个项目所需要的准备时间</div>
                </template>
              </n-form-item>
            </n-gi>
            <n-gi span="1"v-if="formValue.orderMode == 2">
              <n-form-item label="出张通勤时间" path="outWorkTimeDuration" style="margin-bottom: 24px">
                <n-input-group>
                  <n-input-number placeholder="请输入" :min="0" :precision="0" :show-button="false" v-model:value="formValue.outWorkTimeDuration" style="width: 100px" />
                  <n-input-group-label>分钟</n-input-group-label>
                </n-input-group>
                <template #feedback>
                  <div style="font-size: 14px; color: red">当客户预约上门服务，服务技师除了上面的服务间隔时间，还需要赶往上门地址的路上通勤时间</div>
                </template>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预约确定模式" path="orderConfirmType">
                <n-radio-group v-model:value="formValue.orderConfirmType" name="orderConfirmType">
                  <n-space>
                    <n-radio :value="1">
                      手动确认
                    </n-radio>
                    <n-radio :value="2">
                      自动确认
                    </n-radio>
                    <n-radio :value="3">
                      同时支持
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
              <n-form-item label="预计确认时长" path="orderConfirmTime" v-if="formValue.orderConfirmType == 1">
                <n-input-group>
                  <n-input-number placeholder="请输入" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmTime" style="width: 100px" />
                  <n-input-group-label>分钟</n-input-group-label>
                </n-input-group>
              </n-form-item>
              <template v-if="formValue.orderConfirmType == 3 && formValue.orderConfirmDayType == 1">
                <n-form-item label=" " path="orderConfirm">
                  <n-input-group>
                    <n-input-group-label>预定</n-input-group-label>
                    <n-input-number placeholder="请输入" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmBeforeDays" style="width: 100px" />
                    <n-input-group-label>天内手动确认，预计确认时长</n-input-group-label>
                    <n-input-number placeholder="确认时长" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmTime" style="width: 100px" />
                    <n-input-group-label>分钟</n-input-group-label>
                  </n-input-group>
                </n-form-item>
                <n-form-item label=" " path="orderConfirm">
                  <n-input-group>
                    <n-input-group-label>大于{{formValue.orderConfirmBeforeDays}}天自动确认</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </template>
              <template v-if="formValue.orderConfirmType == 3 && formValue.orderConfirmDayType == 2">
                <n-form-item label=" " path="orderConfirm">
                  <n-input-group>
                    <n-input-group-label>预定</n-input-group-label>
                    <n-input-number placeholder="请输入" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmBeforeDays" style="width: 100px" />
                    <n-input-group-label>天前自动确认</n-input-group-label>
                  </n-input-group>
                </n-form-item>
                <n-form-item label=" " path="orderConfirm">
                  <n-input-group>
                    <n-input-group-label>{{formValue.orderConfirmBeforeDays}}天之后手动确认，预计确认时长</n-input-group-label>
                    <n-input-number placeholder="确认时长" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmTime" style="width: 100px" />
                    <n-input-group-label>分钟</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </template>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="提前预约" path="advanceOrderHour" style="margin-bottom: 24px">
                <n-input-group>
                  <n-input-group-label>需至少提前</n-input-group-label>
                  <n-input-number placeholder="请输入小时" :min="0" :precision="0" v-model:value="formValue.advanceOrderHour" style="width: 100px" />
                  <n-input-group-label>小时</n-input-group-label>
                </n-input-group>
                <template #feedback>
                  填0则当前可约
                </template>
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
<!--            <n-gi span="1">-->
<!--              <n-form-item label="派单模式" path="selectStaffType">-->
<!--                <n-radio-group v-model:value="formValue.selectStaffType" name="selectStaffType">-->
<!--                  <n-space>-->
<!--                    <n-radio :value="1">-->
<!--                      人工指定-->
<!--                    </n-radio>-->
<!--                    <n-radio :value="2">-->
<!--                      系统自动指定-->
<!--                    </n-radio>-->
<!--                  </n-space>-->
<!--                </n-radio-group>-->
<!--              </n-form-item>-->
<!--            </n-gi>-->
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="formSubmit">
                确定
              </n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, reactive} from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {MinusCircleOutlined, PlusCircleOutlined} from "@vicons/antd";

  const rules = ref({});
  const group = ref('spaordersetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    orderMode: 2,
    orderTimeType: 1,
    orderTimeWeekArr: [],
    orderTimeWeek: '',
    restTimeWeekArr: [],
    restTimeWeek: '',
    timeDuration: 30,
    serviceTimeDuration: 20,
    outWorkTimeDuration: 20,
    orderConfirmDayType: 1,
    orderConfirmType: 1,
    advanceOrderHour: null,
    orderConfirmBeforeDays: null,
    orderConfirmTime: null,
    orderTimeForm: '',
    advanceOrderDay: null,
    maxOrderDay: null,
    selectStaffType: 1,
  });
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
  const orderTimeFormArr = ref([{ startTime: null, endTime: null }]);
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

  const removeItem = (index: number) => {
    orderTimeFormArr.value.splice(index, 1)
  }

  const addItem = () => {
    orderTimeFormArr.value.push({ startTime: null, endTime: null })
  }

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        formValue.value.orderTimeWeek = formValue.value.orderTimeWeekArr ? formValue.value.orderTimeWeekArr.join(',') : ''
        formValue.value.restTimeWeek = formValue.value.restTimeWeekArr ? formValue.value.restTimeWeekArr.join(',') : ''
        formValue.value.orderTimeForm = JSON.stringify(orderTimeFormArr.value)

        let hasErr = false;
        orderTimeFormArr.value.forEach((item) => {
          if(!item.startTime || !item.endTime){
            message.error('预约时段请填写完整');
            hasErr = true;
            return false;
          }
          // if((new Date('2024-01-01 ' + item.startTime).getTime()) >= (new Date('2024-01-01 ' + item.endTime).getTime())){
          //   message.error('预约时段结束时间应该大于开始时间');
          //   hasErr = true;
          //   return false;
          // }
        })
        if(hasErr){
          formBtnLoading.value = false;
          return false;
        }

        delete formValue.value.orderTimeWeekArr;
        delete formValue.value.restTimeWeekArr;
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          formBtnLoading.value = false;
          message.success('更新成功');
          load();
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      } else {
        message.error('验证失败，请填写完整信息');
        formBtnLoading.value = false;
      }
    });
  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          if(res.list){
            formValue.value = res.list;
            formValue.value.serviceTimeDuration = res.list.serviceTimeDuration ? res.list.serviceTimeDuration : 20
            formValue.value.outWorkTimeDuration = res.list.outWorkTimeDuration ? res.list.outWorkTimeDuration : 20
            let orderTimeWeekArr = res.list.orderTimeWeek ? res.list.orderTimeWeek.split(',') : [];
            formValue.value.orderTimeWeekArr = orderTimeWeekArr.map((item) => {
              return parseInt(item)
            })
            let restTimeWeekArr = res.list.restTimeWeek ? res.list.restTimeWeek.split(',') : []
            formValue.value.restTimeWeekArr = restTimeWeekArr.map((item) => {
              return parseInt(item)
            })
            orderTimeFormArr.value = res.list.orderTimeForm ? JSON.parse(res.list.orderTimeForm) : [{ startTime: null, endTime: null }]
          }
          show.value = false;
        }).catch((err)=>{
          show.value = false;
        })
    });
  }
</script>
