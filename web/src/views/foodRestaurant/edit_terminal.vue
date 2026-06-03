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
          <n-gi>
            <n-form-item label="账号" path="account" :show-require-mark="true">
              <n-input placeholder="请输入账号" v-model:value="formValue.account" :style="{ width: '300px' }" />
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="密码" path="password" :show-require-mark="true">
              <n-input type="password" placeholder="不填则不修改" v-model:value="formValue.password" :style="{ width: '300px' }" />
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="绑定终端" path="terminalType" :show-require-mark="true">
              <n-button type="primary" @click="chooseTerminal('VERIFY_PRINTER')" style="margin-bottom: 25px">绑定核销打印机</n-button>
              <n-button type="primary" @click="chooseTerminal('HAND_TERMINAL')" style="margin-bottom: 25px;margin-left: 10px">绑定手持终端</n-button>
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-table>
              <thead>
              <tr>
                <th>终端编号</th>
                <th>终端名称</th>
                <th>终端类型</th>
                <th>品牌型号</th>
                <th>打印联数</th>
                <th>操作</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="(item, index) of terminalArr" :key="index">
                <td>{{ item.sn }}</td>
                <td>{{ item.terminalName }}</td>
                <td>
                  <n-tag
                    :type="getOptionTag(options.terminal_type, item.terminalType)"
                    size="small"
                    class="min-left-space"
                  >
                    {{ getOptionLabel(options.terminal_type, item.terminalType) }}
                  </n-tag>
                </td>
                <td>
                  <template v-if="item.terminalType === 'VERIFY_PRINTER'">
                    <n-tag
                      :type="getOptionTag(options.verify_brand_model, item.brandModel)"
                      size="small"
                      class="min-left-space"
                    >
                      {{ getOptionLabel(options.verify_brand_model, item.brandModel) }}
                    </n-tag>
                  </template>
                  <template v-else-if="item.terminalType === 'HAND_TERMINAL'">
                    <n-tag
                      :type="getOptionTag(options.hand_brand_model, item.brandModel)"
                      size="small"
                      class="min-left-space"
                    >
                      {{ getOptionLabel(options.hand_brand_model, item.brandModel) }}
                    </n-tag>
                  </template>
                </td>
                <td>
                  <template v-if="item.terminalType == 'VERIFY_PRINTER'">
                    <n-input-number v-model:value="item.printTimes" :min="1" :precision="0" :show-button="false" placeholder="打印联数" style="width: 100px"/>
                  </template>
                  <template v-else>
                    --
                  </template>
                </td>
                <td>
                  <n-button type="error" @click="handleDeleteTerminal(index,item.terminalType,item.id)">删除</n-button>
                </td>
              </tr>
              </tbody>
            </n-table>
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
    <ChooseEditTerminal ref="chooseEditTerminalRef" @reloadTerminal="handleChooseTerminal"/>
  </div>
</template>
<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Edit} from "@/api/foodRestaurant";
import {useMessage} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {useRouter} from "vue-router";
import {getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {options} from "@/views/terminal/model";
import ChooseEditTerminal from "@/components/ChooseTerminal/chooseTerminal.vue";
import {List as terminalList} from "@/api/terminal";

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
  type: 'terminal',
  id: 0,
  account: '',
  password: '',
  terminalIds: '',
  terminalList: [],
  printTerminalIds: '', // 绑定打印机终端ID
  handTerminalIds: '', // 绑定手持终端ID
});
const settingStore = useProjectSettingStore();
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();
const chooseEditTerminalRef = ref();
const chooseTerminalType = ref('')
const terminalArr = ref([])
const terminalIdsArr = ref([])

const rules = ref({});

function chooseTerminal(type){
  chooseTerminalType.value = type
  let chooseTerminalIds = '';
  if(type == 'VERIFY_PRINTER'){
    chooseTerminalIds = formValue.value.printTerminalIds
  }else{
    chooseTerminalIds = formValue.value.handTerminalIds
  }
  chooseEditTerminalRef.value.openModal('RESTAURANT', type, chooseTerminalIds, formValue.value.id);
}

async function getTerminalList(){
  let terminalIds = (formValue.value.printTerminalIds + ',' +  formValue.value.handTerminalIds).replace(/,$/, '')
  const res = await terminalList({
    pagination: false,
    terminalIds: terminalIds,
    needPrintTimesRestaurantId: formValue.value.id,
  });
  if(res.list && res.list.length > 0){
    terminalArr.value = res.list;
    terminalIdsArr.value = res.list.map((item) => {
      return parseInt(item.id)
    })
    formValue.value.terminalIds = terminalIdsArr.value.join(',')
  }else{
    terminalArr.value = []
    terminalIdsArr.value = []
    formValue.value.terminalIds = ''
  }
  show.value = false;
}


function handleDeleteTerminal(index,type,targetId){
  terminalArr.value.splice(index, 1);
  terminalIdsArr.value.splice(index, 1);
  formValue.value.terminalIds = terminalIdsArr.value.join(',')
  if(type == 'VERIFY_PRINTER'){
    formValue.value.printTerminalIds = formValue.value.printTerminalIds.split(',').filter(id => parseInt(id) !== targetId).join(',')
  }else{
    formValue.value.handTerminalIds = formValue.value.handTerminalIds.split(',').filter(id => parseInt(id) !== targetId).join(',')
  }
}

async function handleChooseTerminal(record){
  if(chooseTerminalType.value == 'VERIFY_PRINTER'){
    // 打印机
    formValue.value.printTerminalIds = record
  }else{
    // 手持
    formValue.value.handTerminalIds = record
  }
  show.value = true;
  await getTerminalList()
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      formValue.value.terminalList = terminalArr.value
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
  console.log('data:', data)
  formValue.value = {
    type: 'terminal',
    id: data.id,
    account: data.account,
    password: '',
    terminalIds: '',
    terminalList: [],
    printTerminalIds: data.printTerminalIds, // 绑定打印机终端ID
    handTerminalIds: data.handTerminalIds, // 绑定手持终端ID
  }
  if(data.printTerminalIds || data.handTerminalIds){
    await getTerminalList()
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
