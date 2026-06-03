<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑服务商 #' + formValue.id : '添加服务商' }}</div>
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
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            label-placement="top"
            label-width="auto"
          >
            <div class="statusTab">
              <div :class="newTabValue=='base' ? 'active' : ''"><span @click="handleUpdateValue('base')">基础设置</span></div>
              <div :class="newTabValue=='settlement' ? 'active' : ''"><span @click="handleUpdateValue('settlement')">提成结算设置</span></div>
            </div>
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen" v-if="newTabValue=='base'">
              <n-gi span="1">
                <n-form-item label="服务商名称" path="name">
                  <n-input placeholder="请输入服务商名称" v-model:value="formValue.name" style="width: 300px" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status" :show-require-mark="true">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button
                      v-for="type in options.sys_normal_disable"
                      :key="type.value"
                      :value="type.value"
                      :label="type.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
            </n-grid>
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen" v-if="newTabValue=='settlement'">
              <n-gi span="1">
                <n-form-item label="结算模式" path="settlementId">
                  <n-select
                    v-model:value="formValue.settlementId"
                    :options="settlementList"
                    label-field="name"
                    value-field="id"
                    style="width: 200px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="服务分成" path="settlementType">
                  <n-radio-group v-model:value="formValue.settlementType" name="settlementType">
                    <n-space>
                      <n-radio :value="1">
                        跟随配置
                      </n-radio>
                      <n-radio :value="2">
                        自定义
                      </n-radio>
                    </n-space>
                  </n-radio-group>
                </n-form-item>
                <n-form-item label="分成结算比例" path="orderConfirmHour" v-if="formValue.settlementType == 2">
                  <n-input-group>
                    <n-input-number placeholder="请输入" :show-button="false" :min="0" :max="100" v-model:value="formValue.settlementRate" style="width: 80px" />
                    <n-input-group-label>%</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="服务对象" path="settlementObject">
                  <n-radio-group v-model:value="formValue.settlementObject" name="settlementObject">
                    <n-space>
                      <n-radio value="ISP">
                        服务商
                      </n-radio>
                      <n-radio value="TECHNICIAN">
                        技师
                      </n-radio>
                    </n-space>
                  </n-radio-group>
                </n-form-item>
              </n-gi>
            </n-grid>

          </n-form>

        </n-spin>

      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed} from 'vue';
import {Edit, View} from '@/api/spaIsp';
import {State, newState,loadOptions, options} from './model';
import {List} from "@/api/spaSettlement";
import {useMessage} from "naive-ui";
import {adaModalWidth} from "@/utils/hotgo";


const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const emit = defineEmits(['reloadTable']);
const newTabValue = ref('base')
const formValue = ref<State>(newState(null));

const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();
const settlementList = ref([]);
const rules = ref({
  name: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入服务商名称'
  },
});

function handleUpdateValue(e){
  newTabValue.value = e;
}

async function loadSettlementList(){
  let settlementArrListOrg = await List({
    status: 1,
    Pagination: false
  })
  settlementList.value = settlementArrListOrg.list
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {

      if(!formValue.value.settlementId){
        formBtnLoading.value = false;
        message.error('请选择结算模式');
        newTabValue.value = 'settlement'
        return false;
      }

      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        closeForm();
        emit('reloadTable');
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
}

async function load() {
  let res = await View({ id: formValue.value.id });
  formValue.value = res;
}

async function openModal(state: State) {
  showModal.value = true;
  loading.value = true;
  await loadOptions();
  await loadSettlementList()

  if(state && state.id > 0){
    formValue.value.id = state.id
    await load();
  }else{
    formValue.value = newState(null);
  }
  loading.value = false;
}

function closeForm() {
  showModal.value = false;
  newTabValue.value = 'base'
  loading.value = false;
}

defineExpose({
  openModal,
});

</script>

<style lang="less">
.statusTab{
  display: flex;
  margin-bottom: 30px;
  div{
    margin-right: 5px;
    padding: 0 16px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    span{
      cursor: pointer;
    }
    &.active{
      background: #F2F3F8;
      border-radius: 30px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>


