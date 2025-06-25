# 📘 API Documentation PT Backend ABADI

Base URL: `http://localhost:8080`

---

## 🔐 Authentication

All endpoints **except**:

* `/login`
* `/register`
* `/forgotPassword`
* `/resetPassword`

...require an **Authorization Header**:

```http
Authorization: Bearer <token>
```

Token is valid for **24 hours** after login.

---

## 📄 Auth Endpoints

### ✅ Register

**POST** `/register`

#### Body (x-www-form-urlencoded):

* `name`: string *(required)*
* `email`: string *(required)*
* `password`: string *(required)*
* `role`: `free` or `premium` *(required)*

#### Response:

```json
{
  "message": "User registered"
}
```

---

### 🔐 Login

**POST** `/login`

#### Body (x-www-form-urlencoded):

* `email`: string *(required)*
* `password`: string *(required)*

#### Response:

```json
{
  "token": "<jwt-token-string>"
}
```

---

### 🔁 Forgot Password

**POST** `/forgotPassword`

#### Body (x-www-form-urlencoded):

* `email`: string *(required)*

#### Response:

```json
{
  "message": "Reset token generated. Use this token to reset password.",
  "token": "<reset-token>"
}
```

---

### 🔄 Reset Password

**POST** `/resetPassword`

#### Body (x-www-form-urlencoded):

* `token`: string *(from ForgotPassword)*
* `password`: string *(new password)*

#### Response:

```json
{
  "message": "Password has been reset successfully"
}
```

---

## ❓ Question Endpoints

### ➕ Create Question

**POST** `/questions`
**Headers**: Authorization

#### Body (x-www-form-urlencoded):

* `title`: string *(required)*
* `content`: string *(required)*

---

#### 📥 Get All Questions

**GET** `/questions`
**Headers**: Authorization

---

### 👤 Get Logged-In User's Questions

**GET** `/questions/user`
**Headers**: Authorization *(premium users only)*

---

### 📄 Get Question By ID

**GET** `/questions/{id}`
**Headers**: Authorization

---

### ✏️ Update Question

**PUT** `/questions/{id}`
**Headers**: Authorization

#### Body (x-www-form-urlencoded):

* `title`: string *(optional)*
* `content`: string *(optional)*

---

### ❌ Delete Question

**DELETE** `/questions/{id}`
**Headers**: Authorization

---

## 💬 Answer Endpoints

### ➕ Create Answer

**POST** `/answers`
**Headers**: Authorization

#### Body (x-www-form-urlencoded):

* `questionID`: integer *(required)*
* `content`: string *(required)*

---

### 📥 Get Answers by Question ID

**GET** `/answers/{questionID}`
**Headers**: Authorization

---

### ✏️ Update Answer

**PUT** `/answers/{id}`
**Headers**: Authorization

#### Body (x-www-form-urlencoded):

* `content`: string *(optional)*

---

### ❌ Delete Answer

**DELETE** `/answers/{id}`
**Headers**: Authorization


