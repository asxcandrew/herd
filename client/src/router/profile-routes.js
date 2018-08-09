export default [
  {
    name: 'me',
    path: '/me',
    meta: { requiresAuth: true },
    component: require('@/views/profile/layout').default,
    children: [
      {
        name: 'new-story',
        path: 'new-story',
        component: require('@/views/profile/new-story').default,
      },
      {
        name: 'stories',
        path: 'stories',
        redirect: { name: 'drafts' },
        component: { template: '<router-view></router-view>' },
        children: [
          {
            name: 'drafts',
            path: 'drafts',
            component: require('@/views/profile/stories').default,
          }
        ],
      },
    ],
  },
];
