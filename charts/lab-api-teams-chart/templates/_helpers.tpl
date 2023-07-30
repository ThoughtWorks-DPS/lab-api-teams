{{/*
Expand the name of the chart.
*/}}
{{- define "lab-api-teams.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "lab-api-teams.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "lab-api-teams.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "lab-api-teams.labels" -}}
app: {{ include "lab-api-teams.name" . }}
version: {{ .Values.image.tag }}
env: {{ .Release.Namespace }}
helm.sh/chart: {{ include "lab-api-teams.chart" . }}
{{ include "lab-api-teams.selectorLabels" . }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "lab-api-teams.selectorLabels" -}}
app.kubernetes.io/name: {{ include "lab-api-teams.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "lab-api-teams.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "lab-api-teams.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
