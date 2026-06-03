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
          <n-gi span="1">
            <n-form-item label="出勤日期" path="orderTimeType">
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
            <n-form-item label="定休日" path="restTimeWeekArr">
              <n-checkbox-group v-model:value="formValue.restTimeWeekArr">
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
        </n-grid>
        <div style="text-align: center;margin-top: 30px">
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
import {onMounted, ref, watch} from "vue";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Edit} from "@/api/spaTechnician";
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {useRouter} from "vue-router";

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
const formValue = ref({
  type: 'order',
  id: 0,
  orderTimeType: 1,
  orderTimeWeekArr: [],
  orderTimeWeek: '',
  restTimeWeekArr: [],
  restTimeWeek: '',
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
const settingStore = useProjectSettingStore();
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();

const rules = ref({

});

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(formValue.value.orderTimeType == 2 && formValue.value.orderTimeWeekArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请选择出勤日期');
        return false;
      }
      if(formValue.value.restTimeWeekArr.length <= 0){
        formBtnLoading.value = false;
        message.error('请选择定休日');
        return false;
      }
      formValue.value.orderTimeWeek = formValue.value.orderTimeWeekArr ? formValue.value.orderTimeWeekArr.join(',') : ''
      formValue.value.restTimeWeek = formValue.value.restTimeWeekArr ? formValue.value.restTimeWeekArr.join(',') : ''
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
              router.push({ name: 'spaTechnicianIndex', params: {  } });
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
  let orderTimeWeekArr = data.orderTimeWeek ? data.orderTimeWeek.split(',') : [];
  let restTimeWeekArr = data.restTimeWeek ? data.restTimeWeek.split(',') : []
  formValue.value = {
    type: 'order',
    id: data.id,
    orderTimeType: data.orderTimeType,
    orderTimeWeek: data.orderTimeWeek,
    orderTimeWeekArr: orderTimeWeekArr.map((item) => {
      return parseInt(item)
    }),
    restTimeWeek: data.restTimeWeek,
    restTimeWeekArr: restTimeWeekArr.map((item) => {
      return parseInt(item)
    }),
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
.mapdiv {
  position: absolute;
  top: 15px;
  left: 15px;
  height: 38px;
  text-align: center;
  color: #3d3d3d;
  padding: 8px 15px;
  background-color: #fff;
  box-shadow: 4px 3px 5px #999999a8;
  z-index: 300;
  .tabview {
    color: #333;
    font-size: 14px;
    .showmap {
      color: #053dc8;
      font-weight: 550;
    }
    img {
      width: 18px;
      height: 18px;
      float: left;
      margin-right: 6px;
    }
  }
}
.mapbtm {
  position: absolute;
  z-index: 300;
  bottom: 0;
  text-align: left;
  width: 100%;
  background: #6a7b8c38;
  padding: 5px 10px;
  font-weight: 550;
}
</style>
