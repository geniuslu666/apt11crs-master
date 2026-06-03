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
<!--          <n-gi span="1">
            <n-form-item label="最大席位" path="maxSeat">
              <n-input-group>
                <n-input-number placeholder="请输入最大席位" :min="0" :precision="0" v-model:value="formValue.maxSeat" style="width: 100px" />
                <n-input-group-label>个</n-input-group-label>
              </n-input-group>
            </n-form-item>
          </n-gi>-->
<!--          <n-gi span="1">
            <n-form-item label="预约座位" path="seatNum">
              <n-data-table
                v-model:checked-row-keys="checkedIds"
                :columns="columns"
                :data="seatArr"
                :pagination="false"
                :bordered="false"
                :row-key="(row) => row.id"
                @update:checked-row-keys="onCheckedRow"
                style="width: 400px"
              />
            </n-form-item>
          </n-gi>-->
          <n-gi span="1">
            <n-form-item label="是否允许吸烟" path="canSmoking">
              <n-radio-group v-model:value="formValue.canSmoking" name="canSmoking">
                <n-space>
                  <n-radio :value="2">
                    禁烟
                  </n-radio>
                  <n-radio :value="1">
                    不禁烟
                  </n-radio>
                </n-space>
              </n-radio-group>
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
import {onMounted, ref, watch, h} from "vue";
import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Edit} from "@/api/foodRestaurant";
import {useMessage, NInputGroup, NInputNumber, NInputGroupLabel} from "naive-ui";
import {useTabsViewStore} from "@/store/modules/tabsView";
import {useRouter} from "vue-router";
import {List as seatList} from "@/api/foodSeat";

const props = defineProps({
  formData: {
    type: Object || null,
    default: null,
  },
});
const seatArrList = ref([])
const seatArr = ref([])
const checkedIds = ref([]);
const columns = [
  {
    type: 'selection',
  },
  {
    title: '座位名称',
    key: 'name',
  },
  {
    title: '同时段可预约数量',
    key: 'num',
    render(row, index) {
      return h(
        NInputGroup,
        {},
        [
          h(
            NInputNumber,
            {
              style:{
                width: '100px'
              },
              value: row.num,
              min: 0,
              precision: 0,
              onUpdateValue(v) {
                seatArr.value[index].num = v
              }
            }
          ),
          h(
            NInputGroupLabel,
            {},
            {
              default: () => '人',
            }
          )
        ]
      )
    }
  }
]
const emit = defineEmits(['reloadInfo']);
const router = useRouter();
const tabsViewStore = useTabsViewStore();
const show = ref(false);
const formValue = ref({
  type: 'other',
  id: 0,
  maxSeat: 0,
  seatList: [],
  canSmoking: 1,
});
const settingStore = useProjectSettingStore();
const formBtnLoading = ref(false);
const formRef = ref<any>({});
const message = useMessage();

const rules = ref({});

function onCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      let saveSeatList = [];
      seatArr.value.forEach((item) => {
        if(checkedIds.value.indexOf(item.id) != -1){
          saveSeatList.push({
            seatId: item.id,
            num: item.num
          })
        }
      })
      formValue.value.seatList = saveSeatList;
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

async function loadSeatList(){
  let seatArrListOrg = await seatList({
    status: 1,
    Pagination: false
  })
  seatArrList.value = seatArrListOrg.list
}

async function initData(data){
  formValue.value = {
    type: 'other',
    id: data.id,
    maxSeat: data.maxSeat,
    seatList: [],
    canSmoking: data.canSmoking,
  }
  // 处理座位
  let initCheckedIds = data.seat ? data.seat.map((item) => {
    return item.seatId
  }) : []
  let initCheckedNum = data.seat ? data.seat.map((item) => {
    return item.num
  }) : []

  seatArr.value = []
  seatArrList.value.forEach((item) => {
    let cIndex = initCheckedIds.indexOf(item.id)
    let seatArrItem = {
      id: item.id,
      name: item.seatName,
      num: cIndex == -1 ? 0 : initCheckedNum[cIndex],
    }
    seatArr.value.push(seatArrItem)
  })

  checkedIds.value = initCheckedIds;
}

watch(
  () => props.formData,
  async(value) => {
    await initData(value)
  }
);

onMounted(async() => {
  show.value = true;
  await loadSeatList()
  await initData(props.formData)

  show.value = false;
});

</script>

<style lang="less">
</style>
