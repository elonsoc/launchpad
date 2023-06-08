import { NextResponse } from 'next/server';
import { UserAppInformation } from '../application.d';
import applications from '../data.json';
import { config } from '@/config/Constants';
const BACKEND_URL = config.url.BACKEND_API_URL;

// --------- MOCK
// export async function GET(
// 	request: Request,
// 	{
// 		params,
// 	}: {
// 		params: { id: string };
// 	}
// ): Promise<NextResponse> {
// 	const id = params.id;
// 	let application = applications.filter((app) => app.id === parseInt(id));
// 	return NextResponse.json(application);
// }

// export async function PUT(
// 	request: Request,
// 	{
// 		params,
// 	}: {
// 		params: { id: string };
// 	}
// ): Promise<NextResponse> {
// 	const body: UserAppInformation = await request.json();
// 	const id = params.id;
// 	for (let i = 0; i < applications.length; i++) {
// 		if (applications[i].id.toString() === id) {
// 			applications[i] = { ...applications[i], ...body };
// 		}
// 	}

// 	return NextResponse.json(applications);
// }

// -------------

export async function GET(
	request: Request,
	{
		params,
	}: {
		params: { id: string };
	}
): Promise<NextResponse> {
	const id = params.id;
	let applicationJSON;
	try {
		const res = await fetch(`${BACKEND_URL}/applications/${id}`, {
			headers: {
				'Content-Type': 'application/json',
				credentials: 'include',
			},
		});
		applicationJSON = await res.json();
	} catch (error) {
		console.log('There was an error', error);
		return NextResponse.json({'error': error})
	}
	return NextResponse.json(applicationJSON)
}

export async function PUT(
	request: Request,
	{
		params,
	}: {
		params: { id: string };
	}
): Promise<Response> {
	const body: UserAppInformation = await request.json();
	const id = params.id;
	const res = await fetch(`${BACKEND_URL}/applications/${id}`, {
		method: 'PUT',
		headers: {
			'Content-Type': 'application/json',
			credentials: 'include',
		},
	});
	const applications = await res.json();

	return new Response(applications, {
		status: 200,
		headers: {
			'Access-Control-Allow-Origin': '*',
			'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
			'Access-Control-Allow-Headers': 'Content-Type, Authorization',
		},
	});
}
