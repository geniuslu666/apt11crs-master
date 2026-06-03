<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false"
              :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">房型</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" :content-style="{
                    padding: '0 20px 20px',
                  }"
            class="proCard">
<!--      <BasicForm
        ref="searchFormRef"
        @register="register"
        @submit="reloadTable"
        @reset="reloadTable"
        @keyup.enter="reloadTable"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>-->
      <BasicTable
        ref="actionRef"
        :actionColumn="actionColumn"
        :checked-row-keys="checkedIds"
        :columns="columns"
        :request="loadDataTable"
        :resizeHeightOffset="-10000"
        :row-key="(row) => row.id"
        :scroll-x="scrollX"
        @update:checked-row-keys="handleOnCheckedRow"
      >
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" />
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
import {computed, h, reactive, ref, watch} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {BasicTable, TableAction} from '@/components/Table';
import {useForm} from '@/components/Form/index';
import {usePermission} from '@/hooks/web/usePermission';
import {Delete, List, Status} from '@/api/pmsRoomType';
import {schemas} from './model';
import {adaTableScrollX} from '@/utils/hotgo';
import Edit from './edit.vue';
import View from './view.vue';
import {useRouter} from 'vue-router';
import {useUserStore} from '@/store/modules/user';
import {translang} from "@/utils/smjcomm";

const router = useRouter();
  const dialog = useDialog();
  const message = useMessage();
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const searchFormRef = ref<any>({});
  const editRef = ref();
  const viewRef = ref();
  const checkedIds = ref([]);

const userStore = useUserStore();

  // 表格列
  const columns = [
    // {
    //   title: '封面',
    //   key: 'cover',
    //   align: 'left',
    //   width: -1,
    //   render(row) {
    //     return h(NImage, {
    //       width: 32,
    //       height: 32,
    //       src: row.cover,
    //       fallbackSrc: errorImg,
    //       onError: errorImg,
    //       style: {
    //         width: '32px',
    //         height: '32px',
    //         'max-width': '100%',
    //         'max-height': '100%',
    //       },
    //     });
    //   },
    // },
    // {
    //   title: '图集',
    //   key: 'coverList',
    //   align: 'left',
    //   width: -1,
    //   render(row) {
    //     if (!row.coverList) {
    //       return ``;
    //     }
    //     return row.coverList.map((image) => {
    //       return h(NImage, {
    //         width: 32,
    //         height: 32,
    //         src: image,
    //         onError: errorImg,
    //         style: {
    //           width: '32px',
    //           height: '32px',
    //           'max-width': '100%',
    //           'max-height': '100%',
    //           'margin-left': '2px',
    //         },
    //       });
    //     });
    //   },
    // },
    {
      title: '房型名称',
      key: 'name',
      align: 'left',
      width: 260,
      // render(row){
      //   return namegetlang(row.nameLanguage)
      // }
    },
    // {
    //   title: '最低价格',
    //   key: 'basePrice',
    //   align: 'left',
    //   width: -1,
    // },
    {
      title: '入住时间',
      key: 'checkinAt',
      align: 'left',
      width: 100,
    },
    {
      title: '退房时间',
      key: 'checkoutAt',
      align: 'left',
      width: 100,
    },
    // {
    //   title: '预订方式',
    //   key: 'bookingStyle',
    //   align: 'left',
    //   width: -1,
    // },
    // {
    //   title: '房间风格',
    //   key: 'roomStyle',
    //   align: 'left',
    //   width: -1,
    // },
    {
      title: '占用',
      key: 'occupancy',
      align: 'left',
      width: 100,
    },
    {
      title: '面积',
      key: 'size',
      align: 'left',
      width: 100,
    },
    {
      title: '间数',
      key: 'roomNum',
      align: 'left',
      width: 100,
    },
    // {
    //   title: '浴室',
    //   key: 'bathrooms',
    //   align: 'left',
    //   width: -1,
    // },
    // {
    //   title: '清理费',
    //   key: 'cleaningFee',
    //   align: 'left',
    //   width: -1,
    // },
    {
      title: '创建时间',
      key: 'createAt',
      align: 'left',
      width: 160,
    },
  ];

  const actionColumn = reactive({
    width: 200,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '详情',
            onClick: handleView.bind(null, record),
            auth: ['/pmsRoomType/view'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/pmsRoomType/edit'],
          },
          {
            label: translang('禁用'),
            onClick: handleStatus.bind(null, record),
            ifShow: () => {
              return record.isShow === 1;
            },
            auth: ['/pmsRoomType/status'],
          },
          {
            label: translang('启用'),
            onClick: handleStatus.bind(null, record),
            ifShow: () => {
              return record.isShow === 0;
            },
            auth: ['/pmsRoomType/status'],
          },
          // {
          //   label: '删除',
          //   onClick: handleDelete.bind(null, record),
          //   auth: ['/pmsRoomType/delete'],
          // },
        ],
        // dropDownActions: [
        //   {
        //     label: '查看详情',
        //     key: 'view',
        //     auth: ['/pmsRoomType/view'],
        //   },
        // ],
        // select: (key) => {
        //   if (key === 'view') {
        //     return handleView(record);
        //   }
        // },
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  // 加载表格数据
  const loadDataTable = async (res) => {
    if (router.currentRoute.value.query?.uid) {
      res.puid = router.currentRoute.value.query.uid;
    }else{
      res.puid = userStore.getuserPms.uid;
    }

    return await List({ ...searchFormRef.value?.formModel, ...res });
  };

  // 更新选中的行
  function handleOnCheckedRow(rowKeys) {
    checkedIds.value = rowKeys;
  }

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  watch(
    () => userStore.getuserPms,
    () => {
      if(!router.currentRoute.value.query.uid){
        reloadTable();
      }
    },
    {
      immediate: true,
      deep: true,
    }
  );

  // 添加数据
  function addTable() {
    editRef.value.openModal(null);
  }

  // 编辑数据
  function handleEdit(record: Recordable) {
    editRef.value.openModal(record);
  }

  // 查看详情
  function handleView(record: Recordable) {
    viewRef.value.openModal(record);
  }

  // 单个删除
  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  // 批量删除
  function handleBatchDelete() {
    if (checkedIds.value.length < 1) {
      message.error('请至少选择一项要删除的数据');
      return;
    }

    dialog.warning({
      title: '警告',
      content: '你确定要批量删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          checkedIds.value = [];
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

//打开关闭
function handleStatus(record) {
  let content = record.isShow == 1 ? '确定禁用该房型么？' : '确定开启该房型么？';

  dialog.warning({
    title: '信息',
    content: content,
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      const isShow = record.isShow == 1 ? 0 : 1;
      Status({ id: record.id, isShow: isShow }).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          reloadTable();
        });
      });
    },
  });
}
</script>

<style lang="less" scoped></style>
