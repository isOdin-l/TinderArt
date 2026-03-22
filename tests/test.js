import http from "k6/http";
import { check } from "k6";

export let options = {
  stages: [
    { duration: "30s", target: 200 },
    { duration: "2m", target: 800 },
    { duration: "30s", target: 0 },
  ],
};

const API_BASE = "http://localhost:80";
const PHOTO_FILE = open("./photo.jpg", "b");

function pickOneStyle(styles) {
  const index = Math.floor(Math.random() * styles.length);
  return styles[index];
}

function generateUser(index) {
  const username = `user_${index}_${Math.floor(Math.random() * 10000)}`;
  const lat = 0;
  const lon = 0;
  const favArtStyles = ["cubism", "cyberpunk", "impressionism", "realism"];
  return {
    username,
    name: `Name_${index}`,
    surname: `Surname_${index}`,
    email: `${username}@example.com`,
    password: "password123456",
    description: "test user",
    latitude: lat,
    longitude: lon,
    max_dist_meters: 10000,
    fav_art_styles: pickOneStyle(favArtStyles),
  };
}

function createUser(user) {
  let formData = {
    username: user.username,
    name: user.name,
    surname: user.surname,
    email: user.email,
    password: user.password,
    description: user.description,
    latitude: user.latitude,
    longitude: user.longitude,
    max_dist_meters: user.max_dist_meters,
    photos: http.file(PHOTO_FILE, "photo.jpg"),

    fav_art_styles: user.fav_art_styles,
  };

  return http.post(`${API_BASE}/profile/create`, formData);
}

function loginUser(user) {
  const payload = JSON.stringify({
    username: user.username,
    password: user.password,
  });

  return http.post(`${API_BASE}/auth/sign-in`, payload, {
    headers: { "Content-Type": "application/json" },
  });
}

export default function () {
  const userIndex = __VU;
  const user = generateUser(userIndex);

  // 1. Create user
  let res = createUser(user);
  check(res, { "create profile status 200": (r) => r.status === 200 });

  // 2. Login
  res = loginUser(user);
  check(res, { "login status 200": (r) => r.status === 200 });
  const tokens = res.json();

  const authHeader = {
    headers: {
      Authorization: `Bearer ${tokens.access_token}`,
      "Content-Type": "application/json",
    },
  };

  // 3. Get stack
  res = http.get(`${API_BASE}/profile/protected/stack`, authHeader);
  check(res, { "stack status 200": (r) => r.status === 200 });
  const stack = res.json().matches; // массив userId
  if (stack && Array.isArray(stack)) {
    for (const targetId of stack) {
      // console.log(targetId);
    }
  } else {
    //console.log("stack is not iterable", stack);
  }

  // 4. Get profile
  const getPayload = {
    user_id: stack[0],
  };
  res = http.request(
    "GET",
    `${API_BASE}/profile/protected/get`,
    JSON.stringify(getPayload),
    authHeader,
  );

  check(res, { "get profile status 200": (r) => r.status === 200 });

  // 5. Update
  res = http.patch(
    `${API_BASE}/profile/protected/update`,
    JSON.stringify({ max_dist_meters: 5000 }),
    authHeader,
  );
  check(res, { "update preferences status 200": (r) => r.status === 200 });

  // 6. Swipe
  for (const targetId of stack) {
    const swipePayload = {
      targetId,
      decision: true,
    };
    res = http.post(
      `${API_BASE}/swipe/`,
      JSON.stringify(swipePayload),
      authHeader,
    );

    check(res, { "swipe status 200": (r) => r.status === 200 });
  }

  // 7. Refresh
  const refreshPayload = {
    refresh_token: tokens.refresh_token,
  };

  const headerRefresh = {
    headers: { "Content-Type": "application/json" },
  };
  res = http.post(
    `${API_BASE}/auth/refresh`,
    JSON.stringify(refreshPayload),
    headerRefresh,
  );
  check(res, { "refresh token status 200": (r) => r.status === 200 });

  // 7. Delete
  res = http.del(`${API_BASE}/profile/protected/delete`, null, authHeader);
  check(res, { "delete status 200": (r) => r.status === 200 });
}
