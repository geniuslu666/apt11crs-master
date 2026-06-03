<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择终端</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" @click="confirmForm" style="width: 70px;height: 35px;">
            确认
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
            <template #statusSlot="{ model, field }">
              <n-input v-model:value="model[field]" />
            </template>
          </BasicForm>
          <BasicTable :openChecked="true" ref="actionRef" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow"></BasicTable>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, h, ref} from 'vue';
import {adaModalWidth, adaTableScrollX, getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {BasicTable} from "@/components/Table";
import {options, schemas} from '@/views/terminal/model';
import { List } from '@/api/terminal';
import {BasicForm, useForm} from "@/components/Form";
import {isNullObject} from "@/utils/is";
import {NTag, useMessage} from "naive-ui";

  const emit = defineEmits(['reloadTerminal']);
  const loading = ref(false);
  const showModal = ref(false);
const message = useMessage();
  const dialogWidth = computed(() => {
    return adaModalWidth(850);
  });
  const editBindId = ref(0)
const bindType = ref('')
  const terminalType = ref('')
  const terminalIdsArr = ref([])
const checkedIds = ref([]);


const actionRef = ref();
const searchFormRef = ref<any>({});

  const scrollX = computed(() => {
    return adaTableScrollX(columns, 0);
  });

const [register, {}] = useForm({
  gridProps: { cols: '3' },
  labelWidth: 80,
  schemas,
});

const columns = [
  {
    title: '终端编号',
    key: 'sn',
    align: 'left',
    width: 120,
  },
  {
    title: '终端名称',
    key: 'terminalName',
    align: 'left',
    width: 130,
  },
  {
    title: '终端类型',
    key: 'terminalType',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.terminalType)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.terminal_type, row.terminalType),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.terminal_type, row.terminalType),
        }
      );
    },
  },
  {
    title: '品牌型号',
    key: 'brandModel',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.brandModel)) {
        return ``;
      }
      if (row.terminalType === 'VERIFY_PRINTER') {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.verify_brand_model, row.brandModel),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.verify_brand_model, row.brandModel),
          }
        );
      }else if (row.terminalType === 'HAND_TERMINAL') {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.hand_brand_model, row.brandModel),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.hand_brand_model, row.brandModel),
          }
        );
      }else{
        return ``;
      }
    },
  },
];

    // 加载表格数据
    const loadDataTable = async (res) => {
      res.terminalType = terminalType.value
      if(bindType.value == 'RESTAURANT'){
        // 餐厅
        if(editBindId.value > 0){
          res.terminalRestaurantId = editBindId.value
        }else{
          res.restaurantId = -1
        }
      }else if(bindType.value == 'TH_COUPON'){
        // 礼品券
        if(editBindId.value > 0){
          res.terminalStoreId = editBindId.value
        }else{
          res.storeId = -1
        }
      }
      return await List({ ...searchFormRef.value?.formModel, ...res });
    };

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

  async function openModal(selectBindType, selectType, selectIds, selectBindId) {
    bindType.value = selectBindType
    terminalType.value = selectType;
    terminalIdsArr.value = selectIds ? selectIds.split(',').map((item)=>{
      return parseInt(item)
    }) : [];
    editBindId.value = selectBindId
    checkedIds.value = terminalIdsArr.value
    // await loadOptions()
    showModal.value = true;
  }

  function confirmForm(){
    let checkedIdsStr = '';
    if(checkedIds.value.length <= 0){
      message.error('请选择要绑定的终端');
      return false;
    }
    if(checkedIds.value.length > 0){
      checkedIdsStr = checkedIds.value.join(',')
    }
    emit('reloadTerminal',checkedIdsStr);
    showModal.value = false;
  }


  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
