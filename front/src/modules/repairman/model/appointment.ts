export interface GetAppointments {
  Id: number;
  Username: string;
  Repairman: string;
  DateOfAppointment: string;
  TimeOfAppointment: string;
  Accepted: boolean;
  Price: number;
}

export interface Appointment {
  Username: string;
  Repairman: string;
  DateOfAppointment: string;
  TimeOfAppointment: string;
  Price: number;
}
