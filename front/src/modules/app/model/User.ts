export interface User {
  role: string;
  firstName: string;
  lastName: string;
  emailAddress: string;
  username: string;
  password: string;
  banned: boolean;
}

export interface UpdateUser {
  Id: string;
  FirstName: string;
  LastName: string;
  EmailAddress: string;
}

export interface GetUser {
  FirstName: string;
  LastName: string;
  Username: string;
  EmailAddress: string;
  Role: string;
  Address: string;
  Banned: boolean;
}

export interface GetRepairman {
  FirstName: string;
  LastName: string;
  Username: string;
  EmailAddress: string;
  Role: string;
  Address: string;
  Price: number;
  Occupation: string;
  Review: number;
  ReviewCounter: number;
}

export interface Register {
  FirstName: string;
  LastName: string;
  Username: string;
  Password: string;
  EmailAddress: string;
  Role: string;
  MoneyBalance: number;
  Price: number;
  Occupation: string;
  Address: string;
  Review: number;
  ReviewCounter: number;
  Banned: boolean;
}
