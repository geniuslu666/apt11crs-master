<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择礼品券</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
            <template #statusSlot="{ model, field }">
              <n-input v-model:value="model[field]" />
            </template>
          </BasicForm>
          <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"></BasicTable>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

    <ChooseThCouponDate ref="chooseThCouponDateRef" @closeDrawer="closeForm"/>
  </div>
</template>

<script lang="ts" setup>
import {computed, h, reactive, ref} from 'vue';
import {adaModalWidth, adaTableScrollX} from "@/utils/hotgo";
import {BasicTable, TableAction} from "@/components/Table";
import {BasicForm, useForm} from "@/components/Form";
import {schemas, columns, loadOptions} from '@/views/thCoupon/model';
import {List} from '@/api/thCoupon';
import {useDialog, useMessage} from "naive-ui";
import ChooseThCouponDate from "@/views/pmsMember/chooseThCouponDate.vue";

const dialog = useDialog();
const chooseThCouponDateRef = ref();
const message = useMessage();
  const loading = ref(false);
  const showModal = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(1100);
  });
const actionRef = ref();
const searchFormRef = ref<any>({});

  const actionColumn = reactive({
    width: 90,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '发放',
            onClick: handleChoose.bind(null, record),
          },
        ],
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 100,
    schemas,
  });

    // 加载表格数据
    const loadDataTable = async (res) => {
      res.status = 1;
      return await List({ ...searchFormRef.value?.formModel, ...res });
    };

  const couponMemberId = ref(0);

  function openModal(memberId) {
    loadOptions();
    couponMemberId.value = 0;
    showModal.value = true;
    couponMemberId.value = memberId
  }

  function closeForm(){
    showModal.value = false;
  }

function handleChoose(record: Recordable) {
  chooseThCouponDateRef.value.openModal(couponMemberId.value, record.id);
}

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
