#!/usr/bin/env bash

kubectl create serviceaccount vault-in
kubectl apply -f vault-tokenreview-binding.yaml

SECRET_NAME=$(kubectl -n default get serviceaccount vault-in -o jsonpath='{.secrets[0].name}')

TR_ACCOUNT_TOKEN=$(kubectl -n default get secret ${SECRET_NAME} -o jsonpath='{.data.token}' | base64 --decode)

export VAULT_SA_NAME=$(kubectl get sa vault-in -o jsonpath="{.secrets[*]['name']}")
export SA_CA_CRT=$(kubectl get secret $VAULT_SA_NAME -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo)


vault auth enable kubernetes

vault write auth/kubernetes/config kubernetes_host="https://127.0.0.1:41949" kubernetes_ca_cert="$SA_CA_CRT" token_reviewer_jwt=$TR_ACCOUNT_TOKEN

vault write sys/policy/demo-policy policy=@policy.hcl

vault write auth/kubernetes/role/demo-role \
 bound_service_account_names=vault-in \
 bound_service_account_namespaces=default \
 policies=demo-policy \
 ttl=1h

#DEFAULT_ACCOUNT_TOKEN=$(kubectl get secret $VAULT_SA_NAME -o jsonpath="{.data['token']}" | base64 --decode; echo )

