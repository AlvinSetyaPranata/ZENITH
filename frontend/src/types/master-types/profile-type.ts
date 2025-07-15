type Gender = {
  id: number;
  name: string;
  date_created: string;
};

type City = {
  id: number;
  name: string;
  date_created: string;
};

type Religion = {
  id: number;
  name: string;
  date_created: string;
};

type Country = {
  id: number;
  name: string;
  date_created: string;
};

type Province = {
  id: number;
  name: string;
  created_at: string;
};

type Status = {
  id: number;
  name: string;
  date_created: string;
};

type User = {
  id: number;
  email: string;
  role_id: number;
  date_created: string;
  date_updated: string;
};

type Faculty = {
  id: number;
  name: string;
  created_at: string;
};

type StudyProgram = {
  id: number;
  name: string;
  faculty: any; // or `Faculty | null` if known
  created_at: string;
  updated_at: string;
};

export type StudentType = {
  npm: number | null;
  name: string | null;
  gender_id: Gender | null;
  address: string | null;
  city: City | null;
  religion: Religion | null;
  country: Country | null;
  province: Province | null;
  status: Status | null;
  user: User | null;
  faculty: Faculty | null;
  study_program: StudyProgram | null;
  date_created: string | null;
  date_updated: string | null;
};
