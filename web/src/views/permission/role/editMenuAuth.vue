<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :show-icon="false"
      :mask-closable="false"
      preset="dialog"
      :title="'分配 ' + formValue.name + ' 的菜单权限'"
    >
      <n-spin :show="loading" description="请稍候...">
        <div class="py-3 menu-list" :style="{ maxHeight: '90vh', height: '70vh' }">
          <n-input size="small" v-model:value="pattern" placeholder="输入菜单名称搜索" class="mb-2">
            <template #suffix>
              <n-icon size="18" class="cursor-pointer">
                <SearchOutlined />
              </n-icon>
            </template>
          </n-input>
          <n-tree
            block-line
            checkable
            cascade
            check-on-click
            default-expand-all
            virtual-scroll
            :data="treeData"
            :pattern="pattern"
            :expandedKeys="expandedKeys"
            :checked-keys="checkedKeys"
            style="max-height: 950px; overflow: hidden"
            @update:checked-keys="checkedTree"
            @update:expanded-keys="onExpandedKeys"
          />
        </div>
      </n-spin>
      <template #action>
        <n-space class="mt-6" v-if="showImportSelect">
          <n-input-group>
            <n-tree-select
              size="small"
              placeholder="请选择一个要导入的角色"
              :consistent-menu-width="false"
              clearable
              filterable
              default-expand-all
              :options="editRoleOption"
              key-field="id"
              label-field="name"
              :on-update:value="handleImportSelect"
            />
            <div class="mr-2"></div>
            <n-button ghost @click="showImportSelect = false" size="small"> 取消 </n-button>
          </n-input-group>
        </n-space>

        <n-space class="mt-6 space-group" v-if="!showImportSelect" size="small">
          <n-button ghost @click="showImportSelect = true" size="small"> 导入权限 </n-button>
          <n-button type="info" ghost icon-placement="left" @click="packHandle" size="small">
            全部{{ expandedKeys.length ? '收起' : '展开' }}
          </n-button>
          <n-button type="info" ghost icon-placement="left" @click="checkedAllHandle" size="small">
            全部{{ checkedAll ? '取消' : '选择' }}
          </n-button>

          <n-popconfirm @positive-click="confirmForm">
            <template #trigger>
              <n-button type="primary" :loading="formBtnLoading" size="small">提交</n-button>
            </template>
            你正在修改 {{ formValue.name }} 的菜单权限，确定要提交吗？
          </n-popconfirm>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { GetPermissions, getRoleList, UpdatePermissions } from '@/api/system/role';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { NButton, useMessage } from 'naive-ui';
  import { adaModalWidth, getTreeKeys } from '@/utils/hotgo';
  import { findTreeNode, getAllExpandKeys } from '@/utils';
  import { getMenuList } from '@/api/system/menu';
  import { SearchOutlined } from '@vicons/antd';
  import { State, newState } from '@/views/permission/role/model';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formBtnLoading = ref(false);
  const rawRoleOption = ref<State[]>([]);
  const checkedAll = ref(false);
  const treeData = ref([]);
  const expandedKeys = ref<any[]>([]);
  const checkedKeys = ref<any[]>([]);
  const submitMenuIds = ref<any[]>([]);
  const pattern = ref('');
  const showImportSelect = ref(false);

  const editRoleOption = computed<State[]>(() => {
    if (!rawRoleOption.value) {
      return [];
    }
    const role = findTreeNode(rawRoleOption.value, formValue.value.id, 'id');
    if (role) {
      role.disabled = true;
    }
    return rawRoleOption.value;
  });

  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    const params = {
      id: formValue.value.id,
      menuIds: submitMenuIds.value ?? [],
    };
    UpdatePermissions(params)
      .then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          showModal.value = false;
          emit('reloadTable');
        });
      })
      .finally(() => {
        formBtnLoading.value = false;
      });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  function checkedTree(keys) {
    // 页面显示的选中状态保持原样
    checkedKeys.value = keys;
    // 处理提交用的菜单ID，确保包含必要的父级菜单ID
    const processedKeys = processMenuKeys(keys, treeData.value);
    submitMenuIds.value = processedKeys;
  }

  // 处理菜单选择，确保父级菜单ID也被包含
  function processMenuKeys(selectedKeys: any[], treeData: any[]): any[] {
    const resultKeys = new Set(selectedKeys);

    // 递归函数，用于查找并添加必要的父级菜单ID
    function addNecessaryParentKeys(nodes: any[]) {
      for (const node of nodes) {
        if (node.children && node.children.length > 0) {
          // 检查当前节点的子节点中是否有被选中的
          const hasSelectedChild = hasAnySelectedInSubtree(node.children, selectedKeys);
          if (hasSelectedChild) {
            // 只添加父节点，不添加其他兄弟节点
            resultKeys.add(node.key);
          }
          // 递归处理子节点
          addNecessaryParentKeys(node.children);
        }
      }
    }

    // 检查子树中是否有任何节点被选中（递归检查）
    function hasAnySelectedInSubtree(nodes: any[], selectedKeys: any[]): boolean {
      for (const node of nodes) {
        if (selectedKeys.includes(node.key)) {
          return true;
        }
        if (node.children && node.children.length > 0) {
          if (hasAnySelectedInSubtree(node.children, selectedKeys)) {
            return true;
          }
        }
      }
      return false;
    }

    addNecessaryParentKeys(treeData);
    return Array.from(resultKeys);
  }

  // 从完整的权限数据中提取叶子节点（最终菜单项），用于页面显示
  function extractUserSelectedKeys(allMenuIds: any[], treeData: any[]): any[] {
    const leafKeys = new Set<any>();

    // 递归查找所有被选中的叶子节点
    function findSelectedLeafNodes(nodes: any[]) {
      for (const node of nodes) {
        if (allMenuIds.includes(node.key)) {
          if (!node.children || node.children.length === 0) {
            // 叶子节点且被选中，添加到结果中
            leafKeys.add(node.key);
          } else {
            // 有子节点，继续递归查找
            findSelectedLeafNodes(node.children);
          }
        } else if (node.children && node.children.length > 0) {
          // 父节点未被选中，但可能有子节点被选中
          findSelectedLeafNodes(node.children);
        }
      }
    }

    if (allMenuIds){
      findSelectedLeafNodes(treeData);
    }
    return Array.from(leafKeys);
  }

  function onExpandedKeys(keys) {
    expandedKeys.value = keys;
  }

  function packHandle() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = getAllExpandKeys(treeData) as [];
    }
  }

  function checkedAllHandle() {
    if (!checkedAll.value) {
      const allKeys = getTreeKeys(treeData.value);
      checkedKeys.value = allKeys;
      submitMenuIds.value = allKeys;
      checkedAll.value = true;
    } else {
      checkedKeys.value = [];
      submitMenuIds.value = [];
      checkedAll.value = false;
    }
  }

  function handleImportSelect(key: number) {
    showImportSelect.value = false;
    showModal.value = true;
    getPermissions(key);

    // 默认全部展开
    expandedKeys.value = getAllExpandKeys(treeData);
    message.success('导入成功，提交后生效');
  }

  async function loadMenuList() {
    const res = await getMenuList();
    expandedKeys.value = getAllExpandKeys(res.list) as [];
    treeData.value = res.list;
  }

  async function getPermissions(id: number) {
    checkedKeys.value = [];
    submitMenuIds.value = [];
    checkedAll.value = false;
    const res = await GetPermissions({ id: id });
    // 从完整的权限数据中提取用户实际选择的菜单ID用于页面显示
    checkedKeys.value = extractUserSelectedKeys(res.menuIds, treeData.value);
    // 保存完整的权限数据用于提交
    submitMenuIds.value = res.menuIds;
  }

  async function loadDataList() {
    const res = await getRoleList({ pageSize: 100, page: 1 });
    rawRoleOption.value = res.list;
  }

  async function openModal(record: Recordable) {
    loading.value = true;
    formValue.value = newState(record);
    showModal.value = true;
    await loadMenuList();
    await getPermissions(record.id);
    await loadDataList();
    loading.value = false;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less">
  .space-group {
    margin-left: -8px;
    margin-right: -8px;
  }
</style>
