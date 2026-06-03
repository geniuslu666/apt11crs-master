const allModules = import.meta.glob('./*/index.ts', { eager: true });
const modules = {} as any;
Object.keys(allModules).forEach((path) => {
  const fileName = path.split('/')[1];
  modules[fileName] = allModules[path][fileName] || allModules[path].default || allModules[path];
});

// export default modules
// @ts-ignore
import asyncRoute from './async-route';
// @ts-ignore
import user from './user';
// @ts-ignore
import tabsView from './tabs-view';
// @ts-ignore
import lockscreen from './lockscreen';

export default {
  asyncRoute,
  user,
  tabsView,
  lockscreen,
};
