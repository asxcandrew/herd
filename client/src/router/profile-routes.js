import newStory from '../views/profile/new-story';
import profileLayout from '../views/profile/layout';

export default [
  {
    name: 'me',
    path: 'me',
    meta: { requiresAuth: true },
    component: profileLayout,
    children: [
      {
        name: 'new-story',
        path: 'new-story',
        component: newStory,
      },
    ],
  },
];
