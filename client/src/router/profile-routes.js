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
        meta: { navabrPlugin: 'story' },
        component: require('@/views/profile/new-story').default,
      },
      {
        name: 'edit-story',
        path: 'edit-story/:uid',
        meta: { navabrPlugin: 'story' },
        props: true,
        component: require('@/views/profile/edit-story').default,
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
